package main

import "fmt"

/* A channel is a built-in Go type used to send and receive values between goroutines.
Think of a channel as a pipe through which goroutines can communicate.
Channels help you synchronize goroutines and pass data between them safely.

ch := make(chan int)     // Declare a channel of int type  (unbuffered channel)
make(chan T, n)	   //Channel with capacity n (buffered channel)
ch <- 10     //send 10 into channel
x := <- ch   // receive from channel and assign to x
close(ch)		//Once closed, you cannot send to the channel, but you can still receive until it’s empty.

If goroutines share variables without channels or locks, it can cause race conditions.

*/

func sendData(mssge chan string) {
	mssge <- "Sending data"
}

func main() {
	mssge := make(chan string)
	go sendData(mssge)
	rec_msg := <-mssge
	fmt.Println(rec_msg)
}
