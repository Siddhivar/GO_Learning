package main

import "fmt"

func main() {
	ch := make(chan int, 6)
	for i := 1; i < 5; i++ {
		ch <- i
	}
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}

}
