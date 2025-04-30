package main

import (
	"fmt"
	"time"
)
/*A goroutine is a function that executes concurrently with other functions. 
Think of it as a lightweight thread managed by the Go runtime.*/
func task(id int) {
	fmt.Println("doing task: ", id)
}
func main() {
	for i := 0; i < 10; i++ {
		go task(i)
	}
	time.Sleep(time.Second * 2)
}
