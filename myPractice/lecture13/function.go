package main

import "fmt"

//Simple function

func simpleFunction(){

fmt.Println("Simple function to print a message")

}

//Function with parameters

func add(x int , y int){

sum := 0
sum = x + y
fmt.Println(sum)

}

//function with return type
func addReturn(x int, y int) int{

sum := 0
sum = x + y
return sum

}

//Function with named return values 
func rectangle(l int, b int ) (area int){

var perimeter int
perimeter = 2 * (l + b)
fmt.Println("Perimeter : ", perimeter)

area = l * b
return   //Return statement without specify var name 

}

func rectangleMulti(l int, b int ) (area int, perimeter int){

//var perimeter int
perimeter = 2 * (l + b)
//fmt.Println("Perimeter : ", perimeter)

area = l * b

return   //Return statement without specify var name 

}


func main(){

//calling simple function
fmt.Println("Output of simple function : ")
simpleFunction()

//calling of add function
fmt.Println("Output of simple function : ")
add(20, 40)

//calling addReturn function
fmt.Println("Output of addReturn function :")
x := addReturn(20, 30)
fmt.Println(x)

//calling of rectangle function
//a := rectangle(20, 50)
fmt.Println("Output of rectangle func :")
fmt.Println("Area of rectangle is : ", rectangle(20, 50))

//Calling rectangleMulti function

fmt.Println("Output of rectangleMulti function : ")
fmt.Println(rectangleMulti(30, 10))

01678171501

}