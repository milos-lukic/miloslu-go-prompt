package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}


func statusHandler(w http.ResponseWriter, r *http.Request) {
	res := StatusResponse{Status: "ONLINE"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}


func main() {
	http.HandleFunc("/api/status/", statusHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
