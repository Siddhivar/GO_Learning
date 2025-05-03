package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
		w: A writer that sends data back to the user (like the browser).
		r: The incoming HTTP request, containing info like method (GET/POST), headers, etc.
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //Tells Go to call the provided function when someone accesses '/'.
		fmt.Fprintln(w, "Hello World!") //Writes to the http.ResponseWriter
	})
	fmt.Println("Server is running on https://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // ListenAndServe -> starts an HTTP server on port 8080.
	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
}

