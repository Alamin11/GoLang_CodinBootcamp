package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) (string, string) {
	// Code here or
	m := numbers(in)
	max := m[0]
	min := m[0]
	for _, x := range m {
		if max < x {
			max = x
		}
		if min > x {
			min = x
		}
	}
	a := strconv.Itoa(min)
	b := strconv.Itoa(max)
	//s := []string{a, b}
	return a, b
}
func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}
func main() {
	c, d := HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")
	fmt.Print(c, d)
}
