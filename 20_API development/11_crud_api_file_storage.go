package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// File where users are stored
const dataFile = "users.json"

// Read users from file
func loadUsers() ([]User, error) {
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return []User{}, nil // return empty list if file not found
	}
	var users []User
	err = json.Unmarshal(file, &users)
	return users, err
}

// Save users to file
func saveUsers(users []User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dataFile, data, 0644)
}

// Create new user (POST /users)
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	users, _ := loadUsers()
	users = append(users, newUser)
	saveUsers(users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// Get all users (GET /users)
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := loadUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get one user by ID (GET /users/{id})
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	users, _ := loadUsers()
	for _, u := range users {
		if u.ID == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.NotFound(w, r)
}

// Delete user by ID (DELETE /users/{id})
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	users, _ := loadUsers()
	newUsers := []User{}
	for _, u := range users {
		if u.ID != id {
			newUsers = append(newUsers, u)
		}
	}
	saveUsers(newUsers)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

