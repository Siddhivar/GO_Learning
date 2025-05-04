package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User //No persistent database—data lives only in memory.

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	/*Header: Tells the client to expect JSON.
	Encoder: Converts the users slice to JSON and writes it to the response.
	*/
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {	//Ensures only POST requests are accepted.		
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)	//Decoder: Reads the request body and parses JSON into u.
	if err != nil {
		http.Error(w, "Bad request : invalid json", http.StatusBadRequest)
		return
	}
	users = append(users, u)	//Adds the new user to our slice.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", getUserHandler)
	mux.HandleFunc("/users/create", createUserHandler)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

