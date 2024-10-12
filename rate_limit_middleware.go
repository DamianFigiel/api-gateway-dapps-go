package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	ratelimitPerMinute, err := strconv.ParseInt(os.Getenv("RATE_LIMIT_PER_MINUTE"), 10, 64)
	if err != nil {
		println("Failed to parse RATE_LIMIT_PER_MINUTE, defaulting to 60")
		ratelimitPerMinute = 60
	}
	bucket := ratelimit.NewBucketWithRate(float64(ratelimitPerMinute)/60, ratelimitPerMinute)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if bucket.TakeAvailable(1) == 0 {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
