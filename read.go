package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	type Name struct {
		fname string
		lname string
	}

	var filename string
	nameSlice := make([]Name, 0, 0)

	fmt.Printf("Please enter the name of the text file: ")

	_, err := fmt.Scan(&filename)

	if err != nil {
		fmt.Println("Error: %v", err)
	}
	fmt.Println("Reading file...")

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: %v", err)
	}

	var fname, lname string
	firstNameComplete := false

	for _, v := range file {
		strv := string(v)

		switch strv {
		case " ":
			firstNameComplete = true
		case "\n":
			nameSlice = append(nameSlice, Name{fname: fname, lname: lname})
			fname = ""
			lname = ""
			firstNameComplete = false
		default:
			if firstNameComplete {
				lname = lname + strv
			} else {
				fname = fname + strv
			}
		}

	}

	if fname != "" {
		nameSlice = append(nameSlice, Name{fname: fname, lname: lname})
	}

	if len(nameSlice) > 0 {
		fmt.Printf("Found %d names:\n", len(nameSlice))

		for _, v := range nameSlice {
			fmt.Printf("%s %s\n", v.fname, v.lname)
		}
	} else {
		fmt.Println("No names were found.")
	}

}
