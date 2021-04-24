package main

import (
	"fmt"
	"reflect"
)

func main() {
	//array declaration
	var strArray [3]string
	var intArray [3]int

	//Printing the data type of array using 'reflect' package
	fmt.Println(reflect.TypeOf(strArray).Kind().String())
	fmt.Println(reflect.TypeOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
	fmt.Println(reflect.ValueOf(intArray).Kind())

	// Printing type using default package verbs
	fmt.Printf("%T %T", strArray, intArray)
	//Assigning value to the array
	// arr[0] = "Sam"
	// arr[1] = ""
}
