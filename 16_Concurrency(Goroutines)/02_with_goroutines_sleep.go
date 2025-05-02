package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("Task A")
	go task("Task B")

	time.Sleep(4 * time.Second) // wait for goroutines to finish
}
