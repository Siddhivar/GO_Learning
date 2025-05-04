package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GET request: Here is your data")
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST request: User created")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", GetUserHandler).Methods("GET")
	router.HandleFunc("/user", CreateUserHandler).Methods("POST")
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)

}
