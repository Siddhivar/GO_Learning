package main

import "fmt"

func main() {
	// Using var with type
	var name string = "Siddhi"
	fmt.Println("Name:", name)

	// Using var with type inference
	var age = 22
	fmt.Println("Age:", age)

	// Short variable declaration
	city := "Aligarh"
	fmt.Println("City:", city)

	// Multiple variable declarations
	var x, y, z int = 1, 2, 3
	fmt.Println("x:", x, "y:", y, "z:", z)

	// Zero values
	var emptyString string
	var zeroInt int
	fmt.Println("Empty string:", emptyString)
	fmt.Println("Zero int:", zeroInt)
}
