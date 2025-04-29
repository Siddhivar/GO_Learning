package main

import "fmt"

func main() {
	data := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
	}

	for month, days := range data {
		fmt.Printf("%s has %d days\n", month, days)
	}
}
