package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/juju/ratelimit"
)

var ratelimitPerMinute, _ = strconv.ParseInt(os.Getenv("RATE_LIMIT_PER_MINUTE"), 10, 64)
var bucket = ratelimit.NewBucketWithRate(1.0, ratelimitPerMinute)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if bucket.TakeAvailable(1) == 0 {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
