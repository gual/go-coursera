package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputStr string

	for {
		fmt.Printf("Please enter a string and the characters ‘i’, ‘a’, and ‘n’ will be searched: ")

		_, err := fmt.Scan(&inputStr)

		if err != nil {
			fmt.Println("Error: %v", err)
		}

		lInputStr := strings.ToLower(inputStr)

		if strings.Contains(lInputStr, "a") && strings.HasPrefix(lInputStr, "i") && strings.HasSuffix(lInputStr, "n") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
