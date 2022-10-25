/* Your program should present the user with a prompt, “>”, to indicate that the user
 can type a request. Your program should accept one command at a time from the user,
 print out a response, and print out a new prompt on a new line. Your program should
 continue in this loop forever. Every command from the user must be either a
 “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first
string is “newanimal”. The second string is an arbitrary string which will be the
name of the new animal. The third string is the type of the new animal, either “cow”,
 “bird”, or “snake”.  Your program should process each newanimal command by creating
 the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is
 “query”. The second string is the name of the animal. The third string is the name
 of the information requested about the animal, either “eat”, “move”, or “speak”.
 Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and
Speak(), which take no arguments and return no values. The Eat() method should print
 the animal’s food, the Move() method should print the animal’s locomotion, and the
 Speak() method should print the animal’s spoken sound. Define three types Cow, Bird,
  and Snake. For each of these three types, define methods Eat(), Move(), and Speak()
   so that the types Cow, Bird, and Snake all satisfy the Animal interface.
   When the user creates an animal, create an object of the appropriate type.
   Your program should call the appropriate method when the user issues a query
   command.

Submit your Go program source code. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	getName() string
}

type Cow struct{ name string }
type Snake struct{ name string }
type Bird struct{ name string }

func (c Cow) Eat() {
	fmt.Printf("%s eats grass\n", c.name)
}

func (c Cow) Move() {
	fmt.Printf("%s walks\n", c.name)
}

func (c Cow) Speak() {
	fmt.Printf("%s makes moo\n", c.name)
}

func (s Snake) Eat() {
	fmt.Printf("%s eats a mice\n", s.name)
}

func (s Snake) Move() {
	fmt.Printf("%s slithers\n", s.name)
}

func (s Snake) Speak() {
	fmt.Printf("%s makes hsss\n", s.name)
}

func (b Bird) Eat() {
	fmt.Printf("%s eats worms\n", b.name)
}

func (b Bird) Move() {
	fmt.Printf("%s flies\n", b.name)
}

func (b Bird) Speak() {
	fmt.Printf("%s makes peep\n", b.name)
}
func (c Cow) getName() string {
	return c.name
}

func (s Snake) getName() string {
	return s.name
}

func (b Bird) getName() string {
	return b.name
}

func main() {

	sliceOfAnimals := []Animal{}

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Please type a newanimal or query request or press 'X' to quit: \n > ")
	for input.Scan() { //triggers a new loop cycle whenever there’s new input received from the program’s standard input stream.
		userInput := input.Text()
		fmt.Println("> you typed in:", userInput)
		userInput = strings.ToLower(userInput)
		//check if input is x: in which case you have to break the loop
		if strings.Compare(userInput, "x") == 0 {
			fmt.Println("the program will quit")
			break
		} else {
			inputSlice := strings.Fields(userInput)
			inputQuery := inputSlice[0]
			input1 := inputSlice[1]
			input2 := inputSlice[2]

			switch inputQuery {
			case "newanimal":

				if input2 == "cow" {
					sliceOfAnimals = append(sliceOfAnimals, Cow{name: input1})
					fmt.Println("Created it!")
				} else if input2 == "snake" {
					sliceOfAnimals = append(sliceOfAnimals, Snake{name: input1})
					fmt.Println("Created it!")
				} else if input2 == "bird" {
					sliceOfAnimals = append(sliceOfAnimals, Bird{name: input1})
					fmt.Println("Created it!")
				} else {
					fmt.Println("Invalid request")
					fmt.Println(sliceOfAnimals)
				}
			case "query":

				for _, animal := range sliceOfAnimals {
					if animal.getName() == input1 {
						if input2 == "move" {
							animal.Move()
						} else if input2 == "eat" {
							animal.Eat()
						} else if input2 == "speak" {
							animal.Speak()
						}
					}
				}

			default:
				fmt.Println("Invalid request")

			}
		}
		//finally:
		fmt.Print("Please type in an other animal + information combination or press 'X' to quit: \n > ")
	}

}
