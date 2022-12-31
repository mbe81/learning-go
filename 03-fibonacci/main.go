package main

/*
Write a function that will find the 50th number in the Fibonacci Sequence.

Inspired by: https://www.codecademy.com/resources/blog/20-code-challenges/
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var numberOne, numberTwo, numberThree int

	// First two numbers in the Fibonacci sequence
	numberOne = 0
	numberTwo = 1

	// Calculate the next 48 numbers
	for i := 3; i <= 50; i++ {
		numberThree = numberOne + numberTwo
		numberOne = numberTwo
		numberTwo = numberThree
	}

	// Display the 50th number
	fmt.Println("The 50th number in the Fibonacci sequence is: " + strconv.Itoa(numberThree))

}
