package main

import "fmt"

// DEFER is used to delay the execution of a function until the surrounding function completes.
//defer are executed in LIFO order
func main() {

	defer fmt.Println("Third")
	defer fmt.Println("Second")
	defer fmt.Println("First")
	fmt.Println("Hello")

}
