package main

import (
	"fmt"
	"log"
	"net/http"
)

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

func getNo(w http.ResponseWriter, r *http.Request) {}
