package main

import "fmt"

// FOR is the only loop available in GO
func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", i)
	}
}
