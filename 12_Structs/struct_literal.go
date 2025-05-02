package main

import "fmt"

type Student struct {
	Name string
	Roll int
}

func main() {
	s := Student{Name: "Siddhi", Roll: 101}
	fmt.Println(s)
}
