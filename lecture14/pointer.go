package main

import "fmt"

func update(a *int) int{
fmt.Println(a)
*a = *a + 100
return *a

}

func main(){

var x int //regular variable
var y *int //Special type variable 
x = 10
y = &x

fmt.Println("The assigned value of x is : ", x)
fmt.Println("The address of x is : ", &x)
fmt.Println("The address value/reference value of y is : ", y)
fmt.Println("The dereferenced value of y is : ", *y)

z := update(y)

fmt.Println("Now the updated value is  : ", z)





}