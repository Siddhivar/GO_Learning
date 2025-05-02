package main

import (
	"fmt"
	"time"
)

func countdown() {
	for i := 5; i > 0; i-- {
		fmt.Println("Countdown:", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go countdown()
	fmt.Println("Timer started...")
	time.Sleep(6 * time.Second)
}
