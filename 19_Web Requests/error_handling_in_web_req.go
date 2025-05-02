package main

import (
	"fmt"
	"net/http"
)

func main() {
	_, err := http.Get("http:/www.badURL.com")
	if err != nil {
		fmt.Println("Error handled gracefully: ", err)
	} else {
		fmt.Println("this should not be printed")
	}
}
