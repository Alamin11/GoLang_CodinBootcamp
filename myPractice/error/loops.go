package main

import "fmt"

func main() {
	var i int = 0
	for i <= 1000 {
		if i%3 == 0 && i%7 == 0 && i%9 == 0 {
			fmt.Printf("%v \n", i)
			fmt.Printf("yes %v \n", i)
		}
		i++
	}
}
