package main

import "fmt"

func square(x int) int {
	return x * x
}

//multiple return values
func divide(x, y int) (int, int) {
	return x / y, x % y
}

//named return
func getDetails() (name string, age int) {
	name = "Siddhi"
	age = 21
	return
}

//if return has same datatype
func split(sum int) (x, y int) {
	x = sum - 2
	y = sum * 3
	return
}
func main() {
	fmt.Println("Square of 4 is:", square(4))

	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	n, a := getDetails()
	fmt.Println("Name:", n, "Age:", a)

	fmt.Println(split(4))
}
