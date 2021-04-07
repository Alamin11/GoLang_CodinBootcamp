package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	if (a+b <= c) || (b+c <= a) || (c+a <= b) {
		return false
	}
	return true
}
func main() {
	x := IsTriangle(-1, 2, 3)
	fmt.Print(x)
}
