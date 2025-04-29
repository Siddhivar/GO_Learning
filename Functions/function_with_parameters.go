package main

import "fmt"

func add(a int, b int) {
	fmt.Println("Sum:", a+b)
}
func greet(fname string) {
	fmt.Println("Hello!", fname)

}

func main() {
	add(3, 5)
	greet("Siddhi")
}
