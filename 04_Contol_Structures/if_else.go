package main

import "fmt"

func main() {
	age := 20

	// Basic if
	if age >= 18 {
		fmt.Println("You're an adult.")
	}

	// if, else if
	marks := 85
	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}

	// if-else
	isStudent := true
	if isStudent {
		fmt.Println("Welcome, student!")
	} else {
		fmt.Println("Welcome, guest!")
	}
}
