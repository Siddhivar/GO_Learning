package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Define a route handler for the root URL ("/")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Gorilla Mux!")
	})

	// Start the HTTP server on port 8080 and use the router
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
