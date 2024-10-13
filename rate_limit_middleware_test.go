package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/juju/ratelimit"
)

func TestRateLimitMiddleware(t *testing.T) {
	os.Setenv("RATE_LIMIT_PER_MINUTE", "2")
	ratelimitPerMinute, _ := strconv.ParseInt(os.Getenv("RATE_LIMIT_PER_MINUTE"), 10, 64)
	bucket := ratelimit.NewBucketWithRate(float64(ratelimitPerMinute)/60, ratelimitPerMinute)

	tests := []struct {
		name           string
		expectedStatus int
	}{
		{"First request", http.StatusOK},
		{"Second request", http.StatusOK},
		{"Third request", http.StatusTooManyRequests},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			handler := RateLimitMiddleware(bucket, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			handler.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Code)
			}
		})
	}
}
