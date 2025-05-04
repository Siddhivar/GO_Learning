package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	/*mux.Vars(r) is a function provided by Gorilla Mux that extracts path variables from the URL of an incoming HTTP request.
	It returns a map of strings like -> map[string]string{"name": "Siddhi",}
	*/
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello! %s .", name)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", HelloHandler).Methods("GET") //.Methods("GET") ensures only HTTP GET requests match.
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)

}
