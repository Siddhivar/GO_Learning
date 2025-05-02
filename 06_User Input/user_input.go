package main

import (
	"bufio"
	"fmt"
	"os"
)

/*fmt.Scan() parses inputs without newlines.
fmt.Scanln() stops reading at newline.*/

func main() {
	//for simple input like integers, strings, or float values.
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scan(&name) // Input without spaces
	// fmt.Println("Hello,", name)

	// var a, b int
	// fmt.Print("Enter two numbers: ")
	// fmt.Scan(&a, &b)
	// fmt.Println("Sum is:", a+b)

	//for reading full lines, including spaces.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter full name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello,", name)
}
