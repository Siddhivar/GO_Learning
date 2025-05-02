package main

import (
	"fmt"
	"net/url"
)

func main() {
	buildURL := &url.URL{ //creating a pointer to a url.URL struct.
		Scheme:   "https",
		Host:     "www.google.com",                                              //domain name
		Path:     "/learning",                                                   //after domain
		RawQuery: url.Values{"category": {"Books"}, "sort": {"price"}}.Encode(), //.Encode() converts this into a URL-encoded string:
	}
	fmt.Println("Build URL: ", buildURL.String())
}
