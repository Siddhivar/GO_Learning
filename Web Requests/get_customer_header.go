package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
Imagine you're sending a letter (HTTP request) to an office (a web server). There are two ways to do this:

Method 1: http.Get("https://example.com")
Like dropping a pre-written postcard into a mailbox.
-->It's simple.
-->You can't add custom headers (like extra info).

Method 2: http.NewRequest(...) + client.Do(req)
Like writing a custom letter, putting it in an envelope, and using a courier.
-->You can control what goes in the envelope (headers, method, body).
-->You can say exactly how it should be delivered.
More power and flexibility.
*/
func main() {
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil) //nil means there's no request body
	if err != nil {												//req is a pointer to the http.Request object.
		fmt.Println("error in creating request: ", err)
		return
	}
	req.Header.Set("User-Agent", "GoApp/1.0")	//add custom label like "this is from GoApp version 1.0"

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in sending request: ", err)
		return
	}
	fmt.Println("body: ", string(body))
}
