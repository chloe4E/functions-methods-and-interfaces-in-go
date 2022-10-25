/* Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal: 1) the food that it eats,
 2) its method of locomotion, and 3) the sound it makes when it speaks.
 The following table contains the three animals and their associated data which should be hard-coded into your program.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request,
and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:
food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food,
the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.

Submit your Go program source code. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (p *Animal) Eat() {
	fmt.Println(p.food)
}
func (p *Animal) Move() {
	fmt.Println(p.locomotion)
}
func (p *Animal) Speak() {
	fmt.Println(p.noise)
}

func main() {
	// instantiate the animals:
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	/* 	fmt.Println("Enter an animal (cow, bird, or snake) and a request (eat, move or speak) separated by a space: ")
	   	// var then variable name then variable type
	   	var inputUser string
	   	// Taking input from user
	   	fmt.Scanln(&inputUser) */

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Please type an animal (cow, bird, or snake) and a request (eat, move or speak) separated by a space or press 'X' to quit: \n > ")
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
			inputAnimal := inputSlice[0]
			inputReq := inputSlice[1]
			var currentAnimal Animal
			switch inputAnimal {
			case "cow":
				currentAnimal = cow
			case "bird":
				currentAnimal = bird
			case "snake":
				currentAnimal = snake
			default:
				fmt.Println("requested animal not available")
			}

			switch inputReq {
			case "eat":
				currentAnimal.Eat()
			case "move":
				currentAnimal.Move()
			case "speak":
				currentAnimal.Speak()
			default:
				fmt.Println("requested info not available")
			}

		}
		//finally:
		fmt.Print("Please type in an other animal + information combination or press 'X' to quit: \n > ")
	}

}
