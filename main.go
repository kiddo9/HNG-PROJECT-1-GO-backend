package main

import (
	"HNG_PROJECT_1/controllers"
	"fmt"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Contorl-Allow-Origin", "*")
		w.Header().Set("Access-Contorl-Allow-Header", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	handler := corsMiddleware(mux)

	mux.HandleFunc("/", controllers.GetInternDetails)

	var port = ":8080"
	fmt.Println("server running on port: " + port)
	http.ListenAndServe(port, handler)

}
