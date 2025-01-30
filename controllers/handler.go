package controllers

import (
	"HNG_PROJECT_1/models"
	"encoding/json"
	"net/http"
	"time"
)

// GetUser handles GET requests and returns JSON response
func GetInternDetails(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := models.UserResponse{
		Email:       "paschalelechi0@gmail.com",
		CurrentTime: time.Now().Format(time.RFC3339),
		GitHubURL:   "https://github.com/kiddo9/HNG-PROJECT-1-GO-backend",
	}

	json.NewEncoder(w).Encode(response)
}
