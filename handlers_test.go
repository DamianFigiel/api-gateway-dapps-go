package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchLatestBlock(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"jsonrpc":"2.0","id":1,"result":"0xa"}`)
	}))
	defer mockServer.Close()

	blockNumber, err := fetchLatestBlock(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedBlockNumber := "10"
	if blockNumber != expectedBlockNumber {
		t.Errorf("Expected block number %s, got %s", expectedBlockNumber, blockNumber)
	}
}
