package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	//fmt.Println(<-ch) // This will cause a deadlock if there's no sender

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)

}
