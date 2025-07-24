package main

import (
	"log"
	"net/http"
)

// homeHandler handles the "/" route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home handler called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Welcome to the Go HTTP Server!"}`))
}

// healthHandler handles the "/healthz" route
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

// NewServer creates a new HTTP server mux
func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/healthz", healthHandler)
	return mux
}

func main() {
	port := ":8080"
	log.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, NewServer())
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
