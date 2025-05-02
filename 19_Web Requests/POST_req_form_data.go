package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	formData := url.Values{ //url.Values helps you create key-value pairs like "username=siddhi&password=1234

		"username": {"Siddhi"},
		"password": {"1234"},
	}
	resp, err := http.Post(
		"https://httpbin.org/post",           //URL -> You’re sending the form to this address
		"application/x-www-form-urlencoded",  //Content-Type -> This tells the server that you are sending normal form data
		strings.NewReader(formData.Encode()), //formData.Encode() converts key-values to a string
	)
	if err != nil {
		fmt.Println("Error in posting data: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading response: ", err)
	}
	fmt.Println("body:  ", string(body))

}
