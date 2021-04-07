package main

import "fmt"

func FindOdd(seq []int) int {
	// your code here
	//temp := 0
	count := 0
	res := 0
	for i := 0; i < len(seq); i++ {
		//temp = seq[i]
		for j := 0; j < len(seq); j++ {
			if seq[i] == seq[j] {
				count++
			}
		}
		if count%2 == 1 {
			res = seq[i]
			break
		}
	}
	return res
}
func main() {
	var inputArray = []int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}
	x := FindOdd(inputArray)
	fmt.Print(x)
}
