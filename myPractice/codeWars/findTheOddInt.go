package main

import "fmt"

func FindOdd(seq []int) int {
	// your code here
	dict := make(map[int]int)
	count := 1
	var num, c int

	for _, num = range seq {
		dict[num] = count + 1
		if count%2 == 1 {
			c = num
		}

	}

	return c
}
func main() {
	var inputArray = []int{10, 12, 10, 12, 10}
	x := FindOdd(inputArray)
	fmt.Print(x)
}
