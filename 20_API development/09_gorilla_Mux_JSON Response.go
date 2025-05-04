package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	name := vars["name"]
	user := User{ID: id, Name: name}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}/{name}", UserHandler).Methods("GET")
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
