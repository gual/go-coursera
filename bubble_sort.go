package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputStr string
	intSequence := make([]int, 0, 0)

	fmt.Println("Please enter an integer sequence or enter 's' to sort (after 10 items it will automatically begin sorting).")

	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter an integer or 's': ")
		_, err := fmt.Scan(&inputStr)

		if err != nil {
			fmt.Println("Error: %v", err)
		}

		if strings.ToLower(inputStr) == "s" {
			break
		}

		inputInt, err := strconv.Atoi(inputStr)

		if err != nil {
			fmt.Println("Error: %v", err)
		}

		intSequence = append(intSequence, inputInt)
		fmt.Printf("-- Current sequence length: %d\n\n", len(intSequence))
	}

	fmt.Printf("\nInput sequence:")
	for _, v := range intSequence {
		fmt.Printf(" %d", v)
	}

	BubbleSort(intSequence)

	fmt.Printf("\nSorted sequence:")
	for _, v := range intSequence {
		fmt.Printf(" %d", v)
	}

	fmt.Println("\n\nSo long and thanks for all the fish!")
}

// BubbleSort is used to sort an slice of ints from lower to greater
func BubbleSort(intSlice []int) {
	l := len(intSlice)
	for x := 0; x < l-1; x++ {
		for y := 0; y < l-x-1; y++ {
			if intSlice[y] > intSlice[y+1] {
				Swap(intSlice, y)
			}
		}
	}
}

// Swap 2 adjacent elements
func Swap(intSlice []int, i int) {
	tmp := intSlice[i]
	intSlice[i] = intSlice[i+1]
	intSlice[i+1] = tmp
}
