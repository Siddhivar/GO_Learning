package main

import (
	"fmt"
	"time"
)

func somefunc(num string) {
	fmt.Println(num)

}
func main() {
	go somefunc("3")
	go somefunc("2")
	go somefunc("1")
	time.Sleep(time.Second * 2)
	somefunc("6")
	fmt.Println("heyy")
	somefunc("6")
}
