package main

import (
	"fmt"
	"time"
)

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	go count("A")
	go count("B")
	count("Main")
}
