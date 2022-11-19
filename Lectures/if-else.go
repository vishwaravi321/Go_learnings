package main

import "fmt"

func average(a, b, c int) float32 {
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 is greater than Quiz2!")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz1 is less than Quiz2!")
	} else {
		fmt.Println("Quiz1 and Quiz2 both are equal")
	}

	if average(quiz1, quiz2, int(quiz3)) >= 7 {
		fmt.Println("The mark is acceptable!!!!1")
	} else {
		fmt.Println("The tutor did a terrible job!!!")
	}
}
