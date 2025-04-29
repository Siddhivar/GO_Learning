package main

import "fmt"

/*slices are similar to arrays but more powerful and flexible.
length of a slice can grow or shrink as you see fit

3 ways to create a slice:-
 1. use []datatype{values} format
 2. create slice from array
 3.using the make() function

len()-> no of elements in the slice.
cap()-> no. of elements the slice can grow or shrink to.

*/
func main() {

	nums := []int{10, 20, 30, 40}
	fmt.Println("nums:", nums)
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:", cap(nums))

	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	sli := arr1[2:5]
	fmt.Println("Sli:", sli)
	fmt.Println("Length:", len(sli))
	fmt.Println("Capacity:", cap(sli))

	myslice := make([]int, 5, 5)
	fmt.Println("myslice:", myslice)
	fmt.Println("Length:", len(myslice))
	fmt.Println("Capacity:", cap(myslice))

}
