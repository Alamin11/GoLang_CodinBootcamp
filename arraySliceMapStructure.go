package main

import "fmt"
import "reflect"
func main(){
	newArr := [3]int{1,2,3}

	//Slice declaration
	slice1 := newArr[0 : 2]
	slice2 := append(slice1, 4, 5)
	slice3 := make([]int,4)

	//Copying the element of slice2 to slice3 using built in function
	 copy(slice3, slice2)
	
	//Getting the underlying type of variable type and printing on screen
	t := reflect.TypeOf(slice1).Kind().String() 
	fmt.Println(t)
	
	//printing the type of slice1 variable data type
	fmt.Printf("%T\n", slice1) 

	//Printing all the data of the created slice 
	fmt.Println(slice1, slice2, slice3) 
}