package main

import "fmt"

func sendData(ch chan<- string) {
	ch <- "sending data"
}
func receiveData(ch <-chan string) {
	fmt.Println(<-ch)
}
func main() {
	ch := make(chan string)
	go sendData(ch) //This runs sendData asynchronously in the background.
	receiveData(ch) //This means it runs synchronously—the main() function waits for it to finish.

}
