package main

import "fmt"
import "reflect"
func main(){
	newArr := [3]int{1,2,3}
	t := reflect.TypeOf(newArr).Kind().String()
	fmt.Println(t)
	fmt.Printf("%T", newArr)
	}