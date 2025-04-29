package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		//return 0, errors.New("\nerror->cannot divide by 0") //custom error
		return 0, fmt.Errorf("cannot divide %d by %d ", a, b)
	}
	return a / b, nil
}
func main() {
	// file, err := os.Open("nonexistingfile.csv")
	// if err != nil {
	// 	fmt.Println("error found-> ", err)
	// 	return
	// }
	// defer file.Close()
	// fmt.Println("File opened successfully.")

	fmt.Println(divide(3, 0))
}
