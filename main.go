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
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Request handler
// func handler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		fmt.Fprintln(w, "This is a GET request")
// 	case "POST":
// 		fmt.Fprintln(w, "This is a POST request")
// 	case "PATCH":
// 		fmt.Fprintln(w, "This is a PATCH request")
// 	default:
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 	}
// }

func main() {
	mux := http.NewServeMux()

	// This handles ALL methods, but wrapped with middleware
	mux.HandleFunc("/", controllers.GetInternDetails)

	// Apply CORS middleware
	handler := corsMiddleware(mux)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", handler)
}
