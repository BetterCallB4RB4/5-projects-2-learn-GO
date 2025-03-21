package main

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func applyMiddleware(handlerFunction http.HandlerFunc) http.Handler {
	handler := http.HandlerFunc(handlerFunction)
	return chain(handler, loggingMiddleware, loggingMiddleware1, loggingMiddleware2)
}

func chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("another dummy middleware operation 0")
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("another dummy middleware operation 1")
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("another dummy middleware operation 2")
		next.ServeHTTP(w, r)
	})
}
