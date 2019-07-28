package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputFloat float64

	for {
		fmt.Printf("Please enter a float point number: ")

		_, err := fmt.Scan(&inputFloat)

		if err != nil {
			fmt.Println("Error: %v", err)
		}

		fmt.Println(strconv.FormatFloat(inputFloat, 'f', 0, 64))
	}
}
