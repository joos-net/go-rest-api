package http

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware логируем все запросы
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.Proto)

		next.ServeHTTP(w, r)

		log.Printf("Completed in %v", time.Since(start))
	})
}
