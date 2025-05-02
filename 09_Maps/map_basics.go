package main

import "fmt"

// default value of map is nil
func main() {
	student := map[string]string{
		"Name":  "Siddhi",
		"Study": "Btech",
	}
	//create an empty map- two ways
	var a map[string]int
	var b = make(map[string]int)

	fmt.Println("Name:", student["Name"])
	fmt.Println("Study:", student["Study"])
	fmt.Println(a)
	fmt.Println(b)
}
