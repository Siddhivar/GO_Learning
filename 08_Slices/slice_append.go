package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("Original slice:", nums)

	nums = append(nums, 4)
	fmt.Println("After appending 4:", nums)

	nums = append(nums, 5, 6, 7)
	fmt.Println("After appending 5, 6, 7:", nums)

	moreNums := []int{8, 9}
	nums = append(nums, moreNums...)
	fmt.Println("After appending another slice:", nums)
}
