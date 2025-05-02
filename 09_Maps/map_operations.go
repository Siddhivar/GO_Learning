package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Println(m)

	// update the value
	m["two"] = 3
	fmt.Println(m)

	//delete a value
	delete(m, "two")
	fmt.Println(m)

	//adding a value
	m["three"] = 3
	fmt.Println(m)
}
