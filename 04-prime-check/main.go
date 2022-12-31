package main

/*
Write a function that tests if a number, n, is a prime number.

Inspired by: https://www.codecademy.com/resources/blog/20-code-challenges/
*/

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var numberGiven int

	// read number from console
	fmt.Print("Which number to check? ")
	_, err := fmt.Scanln(&numberGiven)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	// check if it can be didived without remainder
	maxNumberToCheck := int(math.Sqrt(float64(numberGiven)))
	for i := 2; i <= maxNumberToCheck; i++ {
		if numberGiven%i == 0 {
			fmt.Print("Number " + strconv.Itoa(numberGiven) + " is NOT a prime number, it is divisible by " + strconv.Itoa(i) + ".")
			return
		}
	}
	// if it cannot be divieded, it is a prime number
	fmt.Print("Number " + strconv.Itoa(numberGiven) + " is not dvisible, it is a prime number.")

}
