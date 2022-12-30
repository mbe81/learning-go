package main

/*
Write a program that prints the numbers from 1 to 100. But for multiples of
three, print Fizz instead of the number, and multiples of five, print Buzz.
For numbers that are multiples of both three and five, print FizzBuzz.

Inspired by: https://www.codecademy.com/resources/blog/20-code-challenges/
*/

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
