package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

//go:embed reasons.json
var reasonsJSON []byte

var reasons = loadReasons()

type No struct {
	Reason string `json:"reason"`
}

func main() {
	port := "3000"
	router := http.NewServeMux()

	router.HandleFunc("GET /no", getNo)

	fmt.Printf("Server starting on port %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Error occured attempting to start server")
	}
}

func getNo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	randIdx := rand.IntN(len(reasons))
	json.NewEncoder(w).Encode(No{Reason: reasons[randIdx]})
}

func loadReasons() []string {
	var reasons []string
	err := json.Unmarshal(reasonsJSON, &reasons)
	if err != nil {
		log.Fatalf("failed to decode embedded reasons.json: %v", err)
	}

	return reasons
}
