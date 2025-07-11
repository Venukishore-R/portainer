package main

import (
	"log"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
