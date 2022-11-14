package main

import "fmt"

func sum(lhs, rhs int) int {
	return lhs + rhs
}

func main() {
	result := sum(2, 2)
	fmt.Println(result)
}
