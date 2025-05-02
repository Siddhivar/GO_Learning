package main

import (
	"fmt"
	"net/url"
)

func main() {
	searchTerm := "Go Programming & Tips"
	/*The url.QueryEscape() function converts the searchTerm into a URL-encoded format,
	  so that it can be safely placed inside a URL.

	  It replaces unsafe characters:
	  --> 'Space' becomes '+'
	  --> '&' becomes '%26'
	*/
	encoded := url.QueryEscape(searchTerm)
	fmt.Println("Encoded search term:", encoded)

	finalURL := "https://search.com?q=" + encoded
	fmt.Println("Final URL:", finalURL)
}
