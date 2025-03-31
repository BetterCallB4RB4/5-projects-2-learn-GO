package main

import (
	"sync"
	"time"
)

// implementi the token bucket algo
type RateLimiter struct {
	tokens         float64
	maxToken       float64
	refillRate     float64
	lastRefillTime time.Time
	mutex          sync.Mutex
}

// implementi IP limiting
type IPRateLimiter struct {
	limiters map[string]*RateLimiter
	mutex    sync.Mutex
}

// FormatString formats a string with the given prefix and suffix.
//
// prefix: The string to prepend to the input string.
// input: The string to format.
// suffix: The string to append to the input string.
//
// Returns the formatted string.
func newRateLimiter(maxToken float64, refillRate float64) *RateLimiter {
	return &RateLimiter{
		tokens:     maxToken,
		maxToken:   maxToken,
		refillRate: refillRate,
		mutex:      sync.Mutex{},
	}
}

// implement token fill logic
func (rateLimiter *RateLimiter) refillTokens() {
	now := time.Now()
	duration := now.Sub(rateLimiter.lastRefillTime).Seconds()
	tokensToAdd := rateLimiter.refillRate * duration
	rateLimiter.tokens += tokensToAdd
	if rateLimiter.tokens > rateLimiter.maxToken {
		rateLimiter.tokens = rateLimiter.maxToken
	}
	rateLimiter.lastRefillTime = now
}

// implement auth logic base on mutex to avoid singleton write concurrency problem
func (r *RateLimiter) Allow() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.refillTokens()
	if r.tokens >= 1 {
		r.tokens--
		return true
	}
	return false
}

func newIPlimiter() *IPRateLimiter {
	return &IPRateLimiter{
		limiters: make(map[string]*RateLimiter),
	}
}

func (ipLimiter *IPRateLimiter) getLimiter(ip string) *RateLimiter {
	ipLimiter.mutex.Lock()
	defer ipLimiter.mutex.Unlock()

	limiter, exist := ipLimiter.limiters[ip]
	if !exist {
		limiter = newRateLimiter(2, 0.05)
		ipLimiter.limiters[ip] = limiter
	}

	return limiter
}
