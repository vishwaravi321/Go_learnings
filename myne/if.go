package main

import "fmt"

func condition(a, b int) bool {
	return a > b
}

func greet() {
	fmt.Println("Hello My Travelor welcome!!!!")
}

func nongreet() {
	fmt.Println("Who are Nigga??????")
}

func main() {
	if condition(7, 10) {
		greet()
	} else {
		nongreet()
	}
}
