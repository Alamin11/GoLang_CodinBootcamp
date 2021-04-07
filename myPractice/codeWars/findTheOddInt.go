package main

import "fmt"

func FindOdd(seq []int) int {
	// your code here
	//temp := 0
	count := 1
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
	var inputArray = []int{10, 12, 10, 12, 10}
	x := FindOdd(inputArray)
	fmt.Print(x)
}
