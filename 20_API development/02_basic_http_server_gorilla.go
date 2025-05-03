package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() //Create a new router using Gorilla Mux

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // refers to the root path of the server.
		fmt.Fprintln(w, "Hello from Gorilla Mux!")
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)

	/*OUTPUT-> Server running at http://localhost:8080
	it means your Go server started successfully and is running, waiting for incoming HTTP requests. */

}
