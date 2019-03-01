package middlewares

import (
	"log"
	"net/http"
)

// UseCustomLogger uses custom logger middleware for logging requests using provided logger.
func UseCustomLogger(logger *log.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%s %s \n", r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
		})
	}
}

// UseLogger uses standard logger middleware for logging requests.
func UseLogger() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s \n", r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
		})
	}
}
