package main

import "fmt"

func main() {
	//Array declaration
	var arr [5]int
	//Printing the zero values of the declared array
	fmt.Println("Empty array : ", arr)

	//The builtin len returns the length of an array
	fmt.Println("Lenth of arr : ", len(arr))

	//Declaring and assigning value in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

}
