package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/juju/ratelimit"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	ratelimitPerMinute, err := strconv.ParseInt(os.Getenv("RATE_LIMIT_PER_MINUTE"), 10, 64)
	if err != nil {
		println("Failed to parse RATE_LIMIT_PER_MINUTE, defaulting to 60")
		ratelimitPerMinute = 60
	}
	bucket := ratelimit.NewBucketWithRate(float64(ratelimitPerMinute)/60, ratelimitPerMinute)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/blockchain", JWTAuthMiddleware(RateLimitMiddleware(bucket, http.HandlerFunc(BlockchainHandler))))

	println("Server started at :8080")
	http.ListenAndServe(":8080", mux)
}
