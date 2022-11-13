//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var color = "Plum"
	fmt.Println("Fav color :", color)

	birthYear, age := 2001, 21
	fmt.Println("My Birthyear :", birthYear)
	fmt.Println("My age :", age)

	var (
		firstInitial = "V"
		lastInitial  = "A"
	)
	fmt.Println("My First Initial :", firstInitial)
	fmt.Println("My Last Initial :", lastInitial)

	var ageInDays int
	fmt.Println("My age in days:", ageInDays)

	ageInDays = age * 365
	fmt.Println("My age in days:", ageInDays)
}
