package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type BlockchainRequest struct {
	Network string `json:"network"`
}

type BlockchainResponse struct {
	BlockNumber string `json:"block_number"`
}

func BlockchainHandler(w http.ResponseWriter, r *http.Request) {
	var req BlockchainRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var blockNumber string
	switch req.Network {
	case "ethereum":
		blockNumber, err = fetchLatestBlock(os.Getenv("ETHEREUM_ENDPOINT"))
		if err != nil {
			http.Error(w, "Failed to fetch latest block", http.StatusInternalServerError)
			return
		}
	case "polygon":
		blockNumber, err = fetchLatestBlock(os.Getenv("POLYGON_ENDPOINT"))
		if err != nil {
			http.Error(w, "Failed to fetch latest block", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Invalid network", http.StatusBadRequest)
		return
	}

	resp := BlockchainResponse{BlockNumber: blockNumber}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func fetchLatestBlock(rpcURL string) (string, error) {
	requestBody := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}`
	resp, err := http.Post(rpcURL, "application/json", strings.NewReader(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	blockNumberHex := result["result"].(string)
	blockNumber, err := strconv.ParseInt(blockNumberHex, 0, 64)
	if err != nil {
		return "", fmt.Errorf("failed to convert block number from hex to decimal: %v", err)
	}
	return fmt.Sprintf("%d", blockNumber), nil
}
