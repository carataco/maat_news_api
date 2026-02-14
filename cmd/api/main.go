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
	"log"
	"net/http"
	"time"

	"github.com/carataco/maat_news_api/internal/config"
	"github.com/carataco/maat_news_api/internal/server"
)

var cfg config.Config

func init() {

	cfg = config.Config{
		Logger:        1,
		Store:         2,
		RateLimiter:   3,
		Authenticator: 4,
		Port:          5,
	}
}

func main() {
	srv := server.New()

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
