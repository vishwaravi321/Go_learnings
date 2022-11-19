//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import (
	"fmt"
)

func main() {
	//* Use a `switch` statement to print the following:
	//  - "newborn" when age is 0
	//  - "toddler" when age is 1, 2, or 3
	//  - "child" when age is 4 through 12
	//  - "teenager" when age is 13 through 17
	//  - "adult" when age is 18+
	var age int
	fmt.Println("Enter the age:")
	fmt.Scan(&age)
	switch x := age; {
	case x == 0:
		fmt.Println("NewBorn")
	case x < 4:
		fmt.Println("Toddler")
	case x < 13:
		fmt.Println("Child")
	case x < 18:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
}
