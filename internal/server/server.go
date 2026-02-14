package server

/*
Role: HTTP server and routing layer.

Responsibilities:

Creates http.Server instance with proper timeouts.

Registers routes/endpoints (currently just /v1/countries).

Chains middlewares manually (auth, rate limiting, logging).

Starts and shuts down the server gracefully.

Why: Isolates HTTP-specific setup from business logic and makes it easy to replace or extend routing/middleware without touching your handlers or services.
*/

import (
	"net/http"

	"github.com/carataco/maat_news_api/internal/handlers"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}

	s.registerRoutes()
	return s
}

func (s *Server) Handler() http.Handler {
	return s.middleware(s.mux)
}

func (s *Server) registerRoutes() {
	// health check
	s.mux.HandleFunc("/v1/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// resources
	s.mux.HandleFunc("/v1/articles", handlers.GetArticles)
	s.mux.HandleFunc("/v1/authors", handlers.GetAuthors)
}
