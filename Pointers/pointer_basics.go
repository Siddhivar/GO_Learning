package main

import "fmt"

//pointer is a variable whose value is the address of another variable
func main() {
	var a int = 20
	var ip *int
	ip = &a

	var ptr *int //nil pointer

	fmt.Printf("address of a variable:%x\n ", &a)
	fmt.Printf("Address stored in ip variable: %x\n", ip)
	fmt.Println("Value of a: ", a)
	fmt.Println("Value of &a: ", &a)

	fmt.Printf("The value of ptr is : %x\n", ptr)

}
