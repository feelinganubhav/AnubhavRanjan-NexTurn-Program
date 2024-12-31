package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs incoming requests and their responses
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Started %s | %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf(
			"Completed %s | %s | %s | %s | %s",
			r.RemoteAddr,
			time.Now().Format(time.RFC3339),
			r.Method,
			r.URL.Path,
			time.Since(startTime),
		)
	})
}
