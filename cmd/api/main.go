package main

/*
Role: Entry point of your application.

Initializes configuration.

Sets up logger (structured logging with zap).

Creates the HTTP server (via server.New).

Starts the server and handles graceful shutdown signals (SIGINT/SIGTERM).

Why: This keeps all initialization logic centralized and separate from the business/data layers.
*/

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/carataco/maat_news_api/internal/config"
	"github.com/carataco/maat_news_api/internal/server"
	"github.com/joho/godotenv"
)

var cfg config.Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	cfg = config.Config{
		Logger:        1,
		Store:         2,
		RateLimiter:   3,
		Authenticator: 4,
		Port:          5,
	}
}

func LambdaHandler() {
	fmt.Println("Running Lambda Handler")

	srv := server.NewServer()
	srv.RegisterRoutes()

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      srv.Handler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("API listening on :8080")
	log.Fatal(httpServer.ListenAndServe())
}

func main() {
	if os.Getenv("LOCAL") == "true" {
		fmt.Println("Running locally")
		LambdaHandler()
		return
	}

	lambda.Start(LambdaHandler)
}
