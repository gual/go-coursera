package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var inputName string
	var inputAddr string

	dataMap := make(map[string]string)

	fmt.Printf("Please enter a name: ")
	_, err := fmt.Scan(&inputName)
	if err != nil {
		fmt.Println("Error: %v", err)
	}
	dataMap["name"] = inputName

	fmt.Printf("Please enter an address: ")
	_, err = fmt.Scan(&inputAddr)
	if err != nil {
		fmt.Println("Error: %v", err)
	}
	dataMap["address"] = inputAddr

	jsonStr, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	fmt.Println("The result json is: ", string(jsonStr))
}
