package main

import "fmt"

func main() {
	var myName = "Vishwa"
	fmt.Println("My name is", myName)

	var name string = "Klaus"
	fmt.Println("name =", name)

	userName := "mrjindoblu"
	fmt.Println("userName =", userName)

	var sum int
	fmt.Println("The sum is ", sum)

	part1, other := 1,5
	fmt.Println("part1 is", part1,"Other is ", other)

	part2, other := 2,5
	fmt.Println("part2 is", part2,"Other is ", other)

	sum = part1 + part2
	fmt.Println("sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
		)
	fmt.Println("lesson Name is", lessonName)
	fmt.Println("lesson Type is", lessonType)

	word1,word2, _ := "Hello","World","!"
	fmt.Println(word1,word2)
}
