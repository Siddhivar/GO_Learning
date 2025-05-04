package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	fmt.Fprintf(w, "You Searched for %s", query)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/search", SearchHandler).Methods("GET")
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
