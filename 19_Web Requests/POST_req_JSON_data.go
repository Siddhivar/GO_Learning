package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 1: Create a Go map with data
	data := map[string]string{
		"username": "siddhi",
		"password": "1234",
	}
	// 2: Convert Go map to JSON
	jsondata, err := json.Marshal(data)
	for err != nil {
		fmt.Println("Error marshalling json: ", err)
	}
	// 3: Send POST request with JSON data
	resp, err := http.Post(
		"https://httpbin.org/post",
		"applications/json",       // tell server you're sending JSON
		bytes.NewBuffer(jsondata), // JSON data as a readable stream
	)
	for err != nil {
		fmt.Println("Error sending response: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	for err != nil {
		fmt.Println("Error reading response: ", err)
	}
	fmt.Println("Body: ", string(body))

}
