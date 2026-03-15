package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type InfoResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
	Hostname  string `json:"hostname"`
}

func main() {
	port := getEnv("PORT", "8080")
	version := getEnv("APP_VERSION", "1.0.0")
	hostname, _ := os.Hostname()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := InfoResponse{
			Status:    "ok",
			Service:   "go-backend",
			Version:   version,
			Timestamp: time.Now().Format(time.RFC3339),
			Hostname:  hostname,
		}
		_ = json.NewEncoder(w).Encode(resp)
	})

	log.Printf("backend started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
