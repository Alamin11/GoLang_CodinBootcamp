package main

import (
	"fmt"
	"reflect"
)

func main() {
	//array declaration
	var strArray [2]string
	var intArray [3]int

	//Printing the data type of array using 'reflect' package
	fmt.Println(reflect.TypeOf(strArray).Kind().String())
	fmt.Println(reflect.TypeOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
	fmt.Println(reflect.ValueOf(intArray).Kind())

	// Printing type using default package verbs
	fmt.Printf("%T %T\n", strArray, intArray)

	//Printing the length of array
	fmt.Println(len(strArray), len(intArray))
	//Assigning value to the array
	strArray[0] = "Sam"
	strArray[1] = "Micky"

	intArray[0] = 90
	intArray[1] = 110
	intArray[2] = 10
	//Printing aLL THE ARRAY ELEMENT/ITEM
	fmt.Println(strArray)
	fmt.Println(intArray)

	//Accesing the item of array using array name and index
	fmt.Println(intArray[0], strArray[1])

}
