package main

import "fmt"

func main() {
	// Declaring an array of integers
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array elements:", numbers)

	// Shorthand declaration
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits:", fruits)

	// Partial initialization
	languages := [5]string{"Go", "Python"}
	fmt.Println("Languages:", languages)

	// Length of an array
	fmt.Println("Length of fruits array:", len(fruits))

	//initialize only specific elements
	arr1 := [5]int{1: 10, 3: 40}
	fmt.Println(arr1)

	//length is inferred
	arr2 := [...]int{1, 2, 3}
	fmt.Println(len(arr2))

	// Iterating over an array
	for index, value := range numbers {
		fmt.Printf("numbers[%d] = %d\n", index, value)
	}
}
