package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /health", GetHealth)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}
