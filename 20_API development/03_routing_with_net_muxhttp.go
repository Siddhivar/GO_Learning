package main

import (
	"fmt"
	"net/http"
)

/*
An HTTP multiplexer, often just called a mux, is like a traffic director for incoming web requests.
When your server receives a request, it looks at the request’s URL path (e.g., /about or /contact) and asks the mux:

“Which handler should I run for this path?”

The mux then routes the request to the correct handler function you registered.
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page.")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the about page.")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contact us at xyz@gmail.com")
}
func main() {
	mux := http.NewServeMux() //explicitly create a brand–new mux with its own rules.
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/about", AboutHandler)
	mux.HandleFunc("/contact", ContactHandler)

	fmt.Println("Server running on http://localhost:8080")
	//err := http.ListenAndServe(":8080", nil)		that last parameter (nil) means “use the default HTTP multiplexer” instead of mux you created.
	err := http.ListenAndServe(":8080", mux) //ListenAndServe is a blocking call—it starts the server and then waits forever
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
