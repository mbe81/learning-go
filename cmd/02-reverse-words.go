package main

/*
Write a function that will take a given string and reverse the order of the words.

Inspired by: https://www.codecademy.com/resources/blog/20-code-challenges/
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter multiple words to reverse: ")

	// Read line from console and create a slice from it
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	words := strings.Fields(input)

	// Walk through slice in reverse order and print the words
	fmt.Print("Your words in reverse order: ")
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}

}
