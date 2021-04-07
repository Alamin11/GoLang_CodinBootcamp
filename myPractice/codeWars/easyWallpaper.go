package main

import (
	"fmt"
	"math"
)

func WallPaper(l, w, h float64) string {
	// your code
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}

	if l*w*h == 0 {
		return numbers[0]
	} else {
		return numbers[int(math.Ceil((l*h*2+w*h*2)*1.15/5.2))]
	}
}

func main() {
	x := WallPaper(6.3, 5.8, 3.13)
	fmt.Println(x)
}
