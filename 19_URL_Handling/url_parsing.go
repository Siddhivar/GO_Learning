package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Step 1: Start with a URL string
	myURL := "https://example.com:8080/search?q=golang&sort=recent"

	// Step 2: Parse the URL
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	// Step 3: Print individual components
	fmt.Println("Scheme:", parsedURL.Scheme)       // https
	fmt.Println("Host:", parsedURL.Host)           // example.com:8080
	fmt.Println("Hostname:", parsedURL.Hostname()) // example.com
	fmt.Println("Port:", parsedURL.Port())         // 8080
	fmt.Println("Path:", parsedURL.Path)           // /search
	fmt.Println("Raw Query:", parsedURL.RawQuery)  // q=golang&sort=recent

	// Step 4: Get query parameters as a map
	queryParams := parsedURL.Query()
	fmt.Println("Query Parameters:")
	for key, value := range queryParams {
		fmt.Printf("  %s = %s\n", key, value)
	}

	// Step 5: Modify/Add a query parameter
	queryParams.Set("lang", "en")
	parsedURL.RawQuery = queryParams.Encode()

	// Step 6: Print the new URL
	fmt.Println("\nModified URL:", parsedURL.String())
}
