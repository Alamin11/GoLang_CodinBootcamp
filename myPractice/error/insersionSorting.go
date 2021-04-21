package main

import "fmt"

func main() {
	var n = []int{1, 33, 2, 56, 89, 22, 69}
	var i = 1
	length := len(n)
	for i < length {
		j := i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]
			j--
		}
		i++
	}
	fmt.Println(n)
}
