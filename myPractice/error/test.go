package main

import (
	"fmt"
)

func NbYear(p0 int, percent float64, aug int, p int) int {
	// your code
	total := float64(p0)
	var newTotal float64
	var yr int
	// percent := percent/100
	//aug := 50
	for i := 1; ; i++ {
		newTotal += total + total*percent/100.0 + float64(aug)
		total = newTotal
		if total >= float64(p) {
			yr = i + 1
			break
		}
	}
	return yr
}

func main() {

	count := NbYear(1500, 5, 100, 5000)
	fmt.Println(count)
}
