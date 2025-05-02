package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	fmt.Println("Slice:", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
}
