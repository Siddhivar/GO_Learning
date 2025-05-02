package main

import "fmt"

func main() {

	slice := []string{"Go", "Rust", "Python", "Java"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(i, slice[i])
	}

	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
