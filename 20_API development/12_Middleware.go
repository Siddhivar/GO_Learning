package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/*
Think of middleware like layers of security and processing steps before or after your main handler runs.
In real life, imagine you’re entering an office:
First, a security guard checks your ID.
Then a receptionist guides you to the right department.
Finally, you meet the person you actually came to see (your main handler).

Middleware in Go does the same:
Before your handler runs, it can log requests, authenticate users, validate data, etc.
After the handler, it might log responses or modify headers.
*/
func Loggingmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() //Records the start time of the request processing.
		fmt.Println("Request received at: ", start.Format(time.RFC1123))
		fmt.Printf("Method: %s, URL: %s", r.Method, r.URL)
		next.ServeHTTP(w, r) //Passes control to the next handler in the chain
		duration := time.Since(start)
		fmt.Println("Request processed in: ", duration)
	})
}
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}
func main() {
	r := mux.NewRouter()
	r.Use(Loggingmiddleware) //Attach middleware globally
	r.HandleFunc("/", MainHandler).Methods("GET")
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
