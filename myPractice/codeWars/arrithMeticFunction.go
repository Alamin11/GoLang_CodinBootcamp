package main

import "fmt"

func Arithmetic(a int, b int, operator string) int {
	//your code here
	if operator == "add" {
		return a + b
	} else if operator == "subtract" {
		return a - b
	} else if operator == "multiply" {
		return a * b
	} else {
		return a / b
	}
	return 0
}

func main() {
	fmt.Print(Arithmetic(2, 10, "subtract"))
}
