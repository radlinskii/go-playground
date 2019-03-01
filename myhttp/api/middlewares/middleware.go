package middlewares

import (
	"net/http"
)

// Middleware is a wrapper over a standard http.Handler.
type Middleware func(http.Handler) http.Handler

// Apply takes middleware chain and wrappes them around the handler.
func Apply(h http.Handler, middlewares ...Middleware) http.Handler {
	if len(middlewares) == 0 {
		panic("applyMiddlewares used withour a reason")
	}

	l := len(middlewares)
	for i := range middlewares {
		h = middlewares[l-i-1](h)
	}
	return h
}
