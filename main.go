package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
)

var reasons = loadReasons()

type No struct {
	Reason string `json:"reason"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /no", getNo)

	fmt.Println("Server starting on port 3000")
	err := http.ListenAndServe(":3000", router)
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
	data, err := os.ReadFile("./reasons.json")
	if err != nil {
		log.Fatal("An error occurred trying to read reasons.json")
	}

	var reasons []string
	err = json.Unmarshal(data, &reasons)
	if err != nil {
		log.Fatal("An error occured trying to unmarshal the reasons")
	}

	return reasons
}
