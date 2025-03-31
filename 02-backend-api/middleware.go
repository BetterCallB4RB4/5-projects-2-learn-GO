package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strconv"
)

type Middleware func(http.Handler) http.Handler

func applyMiddleware(handlerFunction http.HandlerFunc) http.Handler {
	handler := http.HandlerFunc(handlerFunction)
	return chain(handler, rateLimiterMiddleWare, loggingMiddleware, inputValidation, tokenValidation)
}

func chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// copy the body and pass it through
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error("Error reading request body", "error", err)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var numbers Operation
		err = json.Unmarshal(bodyBytes, &numbers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("couldn't read the request body, cannot perform any operation")
		}

		num1 := strconv.Itoa(numbers.Number1)
		num2 := strconv.Itoa(numbers.Number2)

		logger.Info("Executing operation: " + num1 + " " + path + " " + num2)

		next.ServeHTTP(w, r)
	})
}

func inputValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		processRequest := true

		// check the body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Error reading request body", "error", err)
			processRequest = false
		}

		// restore the message body for other middleware or handler function
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// check if the body respect the structure
		var numbers Operation
		err = json.Unmarshal(bodyBytes, &numbers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("couldn't read the request body, cannot perform any operation")
			processRequest = false
		} else {
			if numbers.Number1 == 0 || numbers.Number2 == 0 {
				http.Error(w, "bad request", http.StatusBadRequest)
				logger.Error("the body is wrongly formatted, cannot perform any operation")
				processRequest = false
			}
		}

		if processRequest {
			next.ServeHTTP(w, r)
		}
	})
}

func tokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !authenticate(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			logger.Warn("protected resources accessed")
			next.ServeHTTP(w, r)
		}
	})
}

func rateLimiterMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Your Ip is rate blocked wait for access this resource", http.StatusInternalServerError)
			return
		}

		limiter := ipRateLimiter.getLimiter(ip)
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
		}
	})
}
