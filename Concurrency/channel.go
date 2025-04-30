package main

import "fmt"

/* A channel is a built-in Go type used to send and receive values between goroutines.
Think of a channel as a pipe through which goroutines can communicate.
Channels help you synchronize goroutines and pass data between them safely.*/

func sendData(ch chan string) {
	ch <- "Hello from goroutine"
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	msg := <-ch
	fmt.Println("Received:", msg)
}
