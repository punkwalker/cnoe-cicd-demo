package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Status       string            `json:"status"`
	Version      string            `json:"version"`
	Env          string            `json:"env"`
	TimeStamp    time.Time         `json:"timestamp"`
	Dependencies map[string]string `json:"dependencies"`
}

// ping handler
func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error : %v", err)
		return
	}

	response := Response{
		Message: "pong from server : " + hostname,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// health handler
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dbConnected := "connected"
	// dbConnected := "not_connected"
	env := os.Getenv("APP_ENV")
	// env := ""

	health := HealthResponse{
		Status:    "healthy",
		Version:   os.Getenv("APP_VERSION"),
		Env:       env,
		TimeStamp: time.Now(),
		Dependencies: map[string]string{
			"database": dbConnected,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/healthz", healthzHandler)
	log.Println("Server started on 8080")
	log.Println("Get a Pong with /ping")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
