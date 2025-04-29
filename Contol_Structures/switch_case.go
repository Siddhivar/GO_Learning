package main

import "fmt"

func main() {
	day := 3
	//It donot need a break statement
	// Basic switch
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

	// Multiple cases
	grade := "B"
	switch grade {
	case "A", "B":
		fmt.Println("Good Job!")
	case "C":
		fmt.Println("Needs Improvement")
	default:
		fmt.Println("Fail")
	}

	// Switch without condition
	num := 45
	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	case num > 0:
		fmt.Println("Positive number")
	}

}
