package main

import (
	"fmt"
	"sync"
	"time"
)

/*Think of it like waiting for your friends to arrive before you go inside a party:

You say: “I'll wait for 3 friends” → wg.Add(3)
Each friend who arrives says: “I'm here!” → wg.Done()
You stand at the door waiting until all 3 arrive → wg.Wait()

Function	What it does
wg.Add(n)	Set how many goroutines to wait for
wg.Done()	Called inside goroutine when it ends
wg.Wait()	Blocks main until all are done
*/

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done() // if remove this line, ERROR -> fatal error: all goroutines are asleep - deadlock!
	for i := 0; i < 6; i++ {
		fmt.Println(name, " : ", i)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go task("taskA", &wg)
	go task("taskB", &wg)
	wg.Wait()
	fmt.Println("all task complete!")
}
