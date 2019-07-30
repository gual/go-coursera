package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface to define its implementations
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	fmt.Println("Welcome. This program will create animals from a predefined set: 'cow', 'bird', 'snake'.")
	fmt.Println("It will also present information ('eat', 'move', 'speak') about the created animals.")
	fmt.Println("In order to create an animal use the following structure 'newanimal <name> <type>'. For example: 'newanimal falco bird'.")
	fmt.Println("In order to obtain information about an animal 'query <name> <info>'. For example: 'query falco speak'.")

	reader := bufio.NewReader(os.Stdin)

	animals := make(map[string]Animal)

	for {
		fmt.Printf("> ")
		inputStr, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error: %v", err)
		}

		s := strings.Split(inputStr, " ")
		name := strings.TrimSpace(s[1])

		switch s[0] {
		case "newanimal":
			var newAnimal Animal

			switch strings.ToLower(strings.TrimSpace(s[2])) {
			case "cow":
				newAnimal = Cow{}
			case "bird":
				newAnimal = Bird{}
			case "snake":
				newAnimal = Snake{}
			default:
				fmt.Println("Error: Animal type not found.")
				continue
			}
			animals[name] = newAnimal
			fmt.Printf("%s was created.\n", name)
		case "query":
			if animal, ok := animals[name]; ok {
				info := strings.ToLower(strings.TrimSpace(s[2]))

				switch info {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("Error: Property not found.")
				}
			} else {
				fmt.Println("Error: Animal not found.")
			}
		default:
			fmt.Println("Error: Command not found.")
			continue
		}
	}
}
