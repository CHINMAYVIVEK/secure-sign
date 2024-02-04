package middleware

import (
	"fmt"
	"net/http"
	"secure-sign/helper"
	"time"
)

// LoggerMiddleware is a middleware for logging incoming requests.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the request details
		duration := time.Since(startTime)
		helper.SugarObj.Info(fmt.Sprintf("%s %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, duration))
		// log.Printf("[%s] %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, duration)
	})
}

// ApplyMiddleware applies multiple middleware functions to the given handler.
// It takes a base handler and a list of middleware functions to be applied sequentially.
// The returned handler will execute each middleware in the order they are provided.
func ApplyMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
