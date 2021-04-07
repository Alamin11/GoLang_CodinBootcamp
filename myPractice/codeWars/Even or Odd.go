package main

import "fmt"

func EvenOrOdd(number int) string {
	// your code here
	//We ca use and operator
	//if number & 1 != 0 {return "Odd"} else{return "even"}
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}
func main() {
	fmt.Print(EvenOrOdd(10019))
}
