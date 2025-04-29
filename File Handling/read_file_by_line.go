package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("error: ", err)
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
