package server

import (
	"log"
	"net/http"
	"time"
)

func (s *Server) middleware(next http.Handler) http.Handler {
	return loggingMiddleware(
		authMiddleware(next),
	)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// placeholder for OAuth
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
