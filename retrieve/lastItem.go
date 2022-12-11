package retrieve

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetLastBid() (int, error) {
	// AWSセッションを作成します
	sess, sessionErr := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")})
	if sessionErr != nil {
		log.Fatal(sessionErr)
	}

	// DynamoDBクライアントを作成します
	svc := dynamodb.New(sess)
	input := &dynamodb.QueryInput{
		TableName: aws.String("shlib_books"),
		ExpressionAttributeNames: map[string]*string{
			"#name": aws.String("name"), // alias付けれたりする
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": { // :を付けるのがセオリーのようです
				S: aws.String("book"),
			},
		},
		KeyConditionExpression: aws.String("#name = :name"), // 検索条件
		ScanIndexForward:       aws.Bool(false),             // ソートキーのソート順（指定しないと昇順）
		Limit:                  aws.Int64(1),                // 取得件数の指定もできる
	}
	var items []Book
	result, err := svc.Query(input)
	if err != nil {
		panic(err)
	}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		panic(err)
	}
	if len(items) == 1 {
		fmt.Println(time.Time(items[0].CreatedAt))
		return items[0].Bid, nil
	} else {
		return 0, err
	}
}
