package main

import "fmt"

func Divisors(n int) int {
	//Put your code here
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(Divisors(5000))
}
