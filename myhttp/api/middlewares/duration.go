package middlewares

import (
	"log"
	"net/http"
	"time"
)

// UseDuration uses duration middleware for logging how long did handlind the request took.
// Probably prefarable as the first element in middleware chain.
func UseDuration(logger *log.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			h.ServeHTTP(w, r)
			logger.Println(time.Now().Sub(t1))
		})
	}
}
