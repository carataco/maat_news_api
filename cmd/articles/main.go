package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/carataco/maat_news_api/internal/articles"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
}

func main() {
	if os.Getenv("LOCAL") == "true" {
		fmt.Println("Running locally")
		return
	}
	lambda.Start(articles.Handler)
}
