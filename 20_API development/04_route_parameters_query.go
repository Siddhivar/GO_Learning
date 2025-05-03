package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()    //Returns all the ?key=value pairs as a map[string][]string.
	name := q.Get("name") //Retrieves the first value for name
	if name == "" {
		name = "Guest"
	}
	lang := q.Get("lang")
	if lang == "" {
		lang = "en"
	}
	fmt.Fprintf(w, "Hello %s, {lang is : %s} ", name, lang)
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	/*
		strings.Trim(r.URL.Path, "/") -> Removes leading/trailing slashes: "/user/42" → "user/42".
		strings.Split(..., "/") -> Splits into []string{"user","42"}.
	*/
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) == 2 && parts[0] == "user" && parts[1] != "" {
		id := parts[1]
		fmt.Fprintf(w, "user id is : %s", id)
		return
	}
	http.NotFound(w, r)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/user/", userHandler)
	fmt.Println("server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	fmt.Println("error starting server: ", err)
}
