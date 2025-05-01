package main

import "fmt"

//pointer is a variable whose value is the address of another variable
func main() {
	// var a int = 20
	// var ip *int
	// ip = &a

	// var ptr *int //nil pointer

	// fmt.Printf("address of a variable:%x\n ", &a)
	// fmt.Printf("Address stored in ip variable: %x\n", ip)
	// fmt.Println("Value of a: ", a)
	// fmt.Println("Value of &a: ", &a)

	// fmt.Printf("The value of ptr is : %x\n", ptr)

	// x := 0xFF
	// y := 0x9C

	// fmt.Printf("type of variable x is %T\n", x)
	// fmt.Printf("type of variable x in hexadecimal is %X\n", x)
	// fmt.Printf("type of variable x in decimal is %v\n", x)

	// fmt.Printf("type of variable x is %T\n", y)
	// fmt.Printf("type of variable x in hexadecimal is %X\n", y)
	// fmt.Printf("type of variable x in decimal is %v\n", y)

	// var x int = 3
	// var p *int
	// p = &x
	// fmt.Println("x= ", x)
	// fmt.Println("p= ", p)
	// fmt.Println("value in p ", *p)
	// fmt.Println("address of x ", &x)

	str := "Golang"
	ptr := &str
	fmt.Println(*ptr)

}

