package main

import "fmt"

// Structures are used to store multiple values of different data types into a single variable
type Person struct {
	Name   string
	Age    int
	Salary int
}

func main() {
	p := Person{"Siddhi", 22, 5000}
	fmt.Println("Name:", p.Name, "Age:", p.Age, "Salary:", p.Salary)
}
