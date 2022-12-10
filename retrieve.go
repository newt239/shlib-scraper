package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

type Book struct {
	bid          int
	isbn         string
	status       string
	localCount   int
	reserveCount int
}

func retrieve() {
	godotenv.Load(".env")
	user_id := os.Getenv("USER_ID")
	password := os.Getenv("PASSWORD")

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
	var bookList []Book
	for i := 1; i < 5; i++ {
		var localCount, reserveCount, status, isbn string
		url := "https://sakaehigashi-lib.com/Library/search/SearchShow.aspx?bid=" + strconv.Itoa(i)
		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.WaitVisible(`#lblLocalCount`),
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
		rep := regexp.MustCompile(`(.)ISBN\n(.)TRC-NO(.)`)
		isbn = rep.ReplaceAllString(isbn, "($2)")
		bookList = append(bookList, Book{
			bid:          i,
			isbn:         isbn,
			status:       status,
			localCount:   localCountInt,
			reserveCount: reserveCountInt,
		})
	}

	for i, v := range bookList {
		fmt.Println(i, v)
	}
	fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
}
