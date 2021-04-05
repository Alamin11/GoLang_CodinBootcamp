package main

import "fmt"

func mul(month int) int {
	x := float32(month)
	if x/4.0 > 0.0 && x/4.0 <= 1.0 {
		return 1
	} else if x/4.0 > 1 && x/4.0 <= 2.0 {
		return 2
	} else {
		return 3
	}
}

func main() {
	res := mul(11)
	fmt.Println(res)
}
