package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func f(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)
// 	}
// }

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt) // Adding some time delay
	}
}

func main() {
	//go f(0)

	//calling n goroutines at a time
	for i := 1; i <= 5; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
