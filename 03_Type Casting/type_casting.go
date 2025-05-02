package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("Value of mean : %f\n", mean)

	//(string(num))-> converts the Unicode code point with value 123 into a string containing that character.
	//Instead use strconv to convert num 123 to "123"
	var num int = 123
	fmt.Println(string(num)) //{

	strr := strconv.Itoa(num)
	fmt.Println(strr)
}
