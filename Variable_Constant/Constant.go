package main

import "fmt"

func main() {
	//value of constant must be assigned when you declare it
	// Constant can be declared both inside and outside of a function
	const PI = 3.14 //written iN uppercase for easy identification
	const COUNTRY string = "India"

	fmt.Println("PI:", PI)
	fmt.Println("Country:", COUNTRY)

	// Constant expressions
	const A int = 5 // Typed Expression
	const a = 5     //Untyped Expression
	const b = a + 10
	fmt.Println("b:", b)
}
