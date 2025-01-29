package main

import (
	"HNG_PROJECT_1/controllers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.GetInternDetails)
	var port = ":8080"
	fmt.Println("server running on port: " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("server not running", err)
	}
}
