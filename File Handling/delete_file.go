package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("data.csv")
	if err != nil {
		fmt.Println("error deleting file: ", err)
		return
	}

	fmt.Println("File deleted successfully.")
}
