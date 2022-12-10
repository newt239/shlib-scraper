package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"Name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Printf("lambda finished! response will be returned!")
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	log.Printf("lambda started!")
	lambda.Start(HandleRequest)
}
