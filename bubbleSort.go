/* Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort()
function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing,
but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Swap(intSlice []int, i int) {
	temp := intSlice[i]
	intSlice[i] = intSlice[i+1]
	intSlice[i+1] = temp
}

func BubbleSort(intSlice []int) {
	l := len(intSlice)
	for j := l - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if intSlice[i] > intSlice[i+1] {
				Swap(intSlice, i)
			}
		}
	}
}

func ReadInputAndConvertToSlice() ([]int, error) {

	// Initiate slice to store user input into
	intSlice := make([]int, 0)

	fmt.Println("Please enter up to 10 integers separated by space")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return intSlice, err
	}

	// Remove the last newline as we used that as a delimiter to read the input
	str = str[:len(str)-1]

	// Convert strings to integers and store them in the slice
	i := 0
	for _, s := range strings.Fields(str) {
		num, err := strconv.Atoi(s)
		if err == nil {
			intSlice = append(intSlice, num)
			// Parse only upto 10 numbers; the rest are ignored
			if i >= 9 {
				break
			}
			i++
		}
	}

	return intSlice, nil

}

func main() {
	// Read input numbers into a slice
	intSlice, err := ReadInputAndConvertToSlice()
	if err != nil {
		log.Fatal(err)
	}

	// Sort the slice
	BubbleSort(intSlice)

	// Print the result
	fmt.Println("\nYour sorted slice is: ", intSlice)

}
