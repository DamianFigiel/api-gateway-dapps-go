package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/blockchain", JWTAuthMiddleware(RateLimitMiddleware(http.HandlerFunc(BlockchainHandler))))

	println("Server started at :8080")
	http.ListenAndServe(":8080", mux)
}
