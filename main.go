package main

import (
	"HNG_PROJECT_1/controllers"
	"fmt"
	"net/http"
)

// CORS Middleware
func main() {
	mux := http.NewServeMux()

	// This handles ALL methods, but wrapped with middleware
	mux.HandleFunc("/", controllers.GetInternDetails)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
