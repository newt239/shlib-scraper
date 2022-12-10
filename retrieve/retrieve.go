package retrieve

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func Retrieve() {
	godotenv.Load(".env")
	user_id := os.Getenv("USER_ID")
	password := os.Getenv("PASSWORD")

	sess, sessionErr := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")})
	if sessionErr != nil {
		log.Fatal(sessionErr)
	}
	db := dynamodb.New(sess)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	start := time.Now()
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://sakaehigashi-lib.com/Library/Login.aspx`),
		chromedp.WaitVisible(`//input[@name="txtLoginUserID"]`),
		chromedp.SendKeys(`//input[@name="txtLoginUserID"]`, user_id),
		chromedp.SendKeys(`//input[@name="txtLoginPassword"]`, password),
		chromedp.Click(`//input[@name="btnLogin"]`, chromedp.NodeVisible),
		chromedp.WaitVisible(`#main-menu`),
		chromedp.Text(`#TopicsControl1_lvTopics_ctrl0_lnkTopic`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 5; i++ {
		var title, author, publish, year, localCount, reserveCount, status, isbn string
		url := "https://sakaehigashi-lib.com/Library/search/SearchShow.aspx?bid=" + strconv.Itoa(i)
		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.WaitVisible(`#lblTitle`),
			chromedp.WaitVisible(`#lblLocalCount`),
			chromedp.Text(`#lblTitle`, &title, chromedp.NodeVisible, chromedp.ByQuery),
			chromedp.Text(`#lblAuthor`, &author, chromedp.NodeVisible, chromedp.ByQuery),
			chromedp.Text(`#lblPublish`, &publish, chromedp.NodeVisible, chromedp.ByQuery),
			chromedp.Text(`#lblYear`, &year, chromedp.NodeVisible, chromedp.ByQuery),
			chromedp.Text(`#lblLocalCount`, &localCount, chromedp.NodeVisible, chromedp.ByQuery),
			chromedp.Text(`#lblRsvCount`, &reserveCount, chromedp.NodeVisible, chromedp.ByQuery),
			chromedp.Text(`#lvSearchLocalList_ctrl0_lblStatus`, &status, chromedp.NodeVisible, chromedp.ByQuery),
			chromedp.Text(".show-xml", &isbn, chromedp.ByQueryAll),
		)
		if err != nil {
			log.Fatal(err)
		}
		if status == "" {
			break
		}
		localCountInt, _ := strconv.Atoi(localCount)
		reserveCountInt, _ := strconv.Atoi(reserveCount)
		rep := regexp.MustCompile(`(.*)ISBN\n(.*)\n(.*)`)
		var book Book
		var info BookInfo
		var data OpenBDResponse
		if rep.MatchString(isbn) {
			isbn = rep.FindStringSubmatch(isbn)[2]
			info, data, err = GetBookInfo(isbn)
			if err != nil {
				log.Fatal(err)
			}
			book = Book{
				Name:         "book",
				Bid:          i,
				Title:        title,
				Author:       longerStr(author, info.Author),
				Publisher:    longerStr(publish, info.Publisher),
				Pubdate:      longerStr(year, info.Pubdate),
				Lanove:       info.Lanove,
				Tameshiyomi:  info.Tameshiyomi,
				Isbn:         isbn,
				Status:       status,
				LocalCount:   localCountInt,
				ReserveCount: reserveCountInt,
				Source:       info.Source,
				CreatedAt:    dynamodbattribute.UnixTime(time.Now()),
			}
		} else {
			book = Book{
				Name:         "book",
				Bid:          i,
				Title:        title,
				Author:       author,
				Publisher:    publish,
				Pubdate:      year,
				Isbn:         "original",
				Status:       status,
				LocalCount:   localCountInt,
				ReserveCount: reserveCountInt,
				Source:       "shlib",
				CreatedAt:    dynamodbattribute.UnixTime(time.Now()),
			}
		}
		inputAV, err := dynamodbattribute.MarshalMap(book)
		if err != nil {
			log.Fatal(err)
		}
		input := &dynamodb.PutItemInput{
			TableName: aws.String("shlib_books"),
			Item:      inputAV,
		}
		_, err = db.PutItem(input)
		if err != nil {
			log.Fatal(err)
		}

		wordList, err := getWordList(book, data)
		if err != nil {
			log.Fatal(err)
		}
		inputAV, err = dynamodbattribute.MarshalMap(BookWord{
			Name: "book",
			Bid:  book.Bid,
			Text: wordList,
		})
		if err != nil {
			log.Fatal(err)
		}
		input = &dynamodb.PutItemInput{
			TableName: aws.String("shlib_bookwords"),
			Item:      inputAV,
		}
		_, err = db.PutItem(input)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
}

func longerStr(t1 string, t2 string) string {
	if len(t1) >= len(t2) {
		return t1
	} else {
		return t2
	}
}
