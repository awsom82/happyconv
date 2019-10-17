package main

import (
	"errors"
	"net/http"

	"golang.org/x/time/rate"
)

var (
	// ErrLimitExceed error for rate limit
	ErrLimitExceed = errors.New("webconv: rate limit exceeded")
)

var limiter = rate.NewLimiter(2, 2)

// rateLimit recived http.Handler and return ErrLimitExceed if limit was hit
func rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, ErrLimitExceed.Error(), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
