package main

import "fmt"

func main() {

	//while like loop
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Printf("\n")
	// for loop with condition
	for j := 1; j <= 10; j++ {
		fmt.Printf("%d ", j)
	}
	fmt.Printf("\n")
	//for loop without condition
	for {
		fmt.Println("Infinite Loop")
		break
	}
	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
