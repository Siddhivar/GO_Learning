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

var users []User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	users = append(users, u)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			var updatedUser User
			json.NewDecoder(r.Body).Decode(&updatedUser)
			users[index] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.NotFound(w, r)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)

}
