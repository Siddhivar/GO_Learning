package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))

	copied := copy(dst, src)
	fmt.Println("Source slice:", src)
	fmt.Println("Destination slice after copy:", dst)
	fmt.Println("Number of elements copied:", copied)
}
