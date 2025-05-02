package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
A web request (like http.Get) means you're asking a web server for some information
(like a webpage, API response, file, etc.).
*/
func main() {
	resp, err := http.Get("https://example.com") //Makes a GET request to the provided URL.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() //It closes the connection after the function ends, otherwise resource leakage.

	body, err := io.ReadAll(resp.Body) //This reads the body of the response into bytes.
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response body:") //Converts bytes to a string so it can be printed.
	fmt.Println(string(body))     // will print the HTML of example.com
}
