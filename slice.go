package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var inputStr string
	intSlice := make([]int, 0, 3)

	for {
		fmt.Printf("Please enter an integer to be added to the slice: ")

		_, err := fmt.Scan(&inputStr)

		if err != nil {
			fmt.Println("Error: %v", err)
		}

		if strings.ToLower(inputStr) == "x" {
			fmt.Println("So long and thanks for all the fish!")
			break
		}

		inputInt, err := strconv.Atoi(inputStr)

		if err != nil {
			fmt.Println("Error: %v", err)
		}

		intSlice = append(intSlice, inputInt)
		sort.Ints(intSlice)

		fmt.Printf("Slice order with new integer included: ")

		for _, v := range intSlice {
			fmt.Printf("%d ", v)
		}

		fmt.Println()
	}
}
