package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}
func sub(a, b, c int) int {
	return add(a, b) - c
}
func greet() {
	fmt.Println("Hello !! from greet function")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A Baker's dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Another Baker's dozeń is ", anotherBakersDozen)
	fmt.Println("Sub function:", sub(5, 5, 5))
}
