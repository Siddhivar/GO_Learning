package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://www.google.com/search?q=golang+url+parsing&hl=en"
	parsed_URL, _ := url.Parse(rawURL)

	fmt.Println("Full URL: ", parsed_URL.String())
	fmt.Println("Scheme: ", parsed_URL.Scheme)
	fmt.Println("Host: ", parsed_URL.Host)
	fmt.Println("Hostname: ", parsed_URL.Hostname())
	fmt.Println("Port: ", parsed_URL.Port())
	fmt.Println("Path: ", parsed_URL.Path)
	fmt.Println("raw query: ", parsed_URL.RawQuery)
	fmt.Println("Raw path: ", parsed_URL.RawPath)

	queryparams := parsed_URL.Query()
	fmt.Println("query parameters")
	for key, val := range queryparams {
		fmt.Printf("%s : %s\n", key, val)
	}

}
