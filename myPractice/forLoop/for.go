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
			continue //force to the next iteration
		}
		fmt.Println(n)
	}

	//Range for loop
	// Example 1
	strDict := map[string]string{"Bangladesh": "Dhaka", "Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}

	//range loop over string
	for range "Hello" {
		fmt.Println("Hello")
	}

}
