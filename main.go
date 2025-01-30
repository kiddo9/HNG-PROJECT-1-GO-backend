package main

import (
	"HNG_PROJECT_1/controllers"
	"fmt"
	"net/http"
)

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")

		// Handle OPTIONS request (Preflight check)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// Register routes **before** applying middleware
	mux.HandleFunc("/", controllers.GetInternDetails)

	// Wrap `mux` with CORS middleware after setting routes
	handler := corsMiddleware(mux)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", handler)
}
