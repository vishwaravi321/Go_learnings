//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) string {
	f := "hello" + name
	return f
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func message() string {
	m := "hello MF!!!!"
	return m
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func number(l int) int {
	return l
}

// * Write a function that returns any two numbers
func twoNumbers(x, y int) (int, int) {
	return x, y
}

// * Add three numbers together using any combination of the existing functions.
//   - Print the result
func combination() {
	q := add(number(5), number(9), number(6))
	fmt.Println("The result is :", q)
}

// * Call every function at least once
func main() {
	fmt.Println("The greet function:", greet("Vishwa"))
	fmt.Println("The message function:", message())
	fmt.Println("Just add function:", add(5, 6, 7))
	fmt.Println("The number Function:", number(9))
	fmt.Println("the Any two numbers function---------")
	u, i := twoNumbers(3, 44)
	fmt.Println("the result of two numbers function is ", u, i)
	fmt.Println("The Combination Function")
	combination()
}
