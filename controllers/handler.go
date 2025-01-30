package controllers

import (
	"HNG_PROJECT_1/models"
	"encoding/json"
	"net/http"
	"time"
)

// GetUser handles GET requests and returns JSON response
func GetInternDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all domains (or specify allowed domains)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	response := models.UserResponse{
		Email:       "paschalelechi0@gmail.com",
		CurrentTime: time.Now().Format(time.RFC3339),
		GitHubURL:   "https://github.com/kiddo9/HNG-PROJECT-1-GO-backend",
	}

	json.NewEncoder(w).Encode(response)
}
