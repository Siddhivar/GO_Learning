package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Request timed out or failed: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in receiving data: ", err)
		return
	}
	fmt.Println("Body:  ", string(body))

}
