package main

import "fmt"

func main() {
	var name string
	var num1 int
	var num2 float64
	var chr rune
	var isBool bool
	//fmt.Println("  | :", name)
	fmt.Println("Zero values of all primitive data types: ")
	fmt.Println("Zero value of string : ", name)
	fmt.Println("Zero value of int : ", num1)
	fmt.Println("Zero value of float : ", num2)
	fmt.Println("Zero value of rune : ", chr)
	fmt.Println("Zero value of bool : ", isBool)

	//Different boolean operation
	fmt.Printf("Boolean Operation:\n")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println(!true)
}
