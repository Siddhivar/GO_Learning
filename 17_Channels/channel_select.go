package main

import (
	"fmt"
	"time"
)

/*
The select statement in Go is used to wait on multiple channel operations—
whichever is ready first gets executed. It’s like a switch, but for channels

Wait on multiple channels without blocking on a specific one.
Handle whichever channel is ready first.
Prevent deadlocks when dealing with multiple concurrent inputs/outputs.
*/
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "One"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received: ", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received: ", msg2)
		default:
			fmt.Println("No messages yet.")
		}
	}

}
