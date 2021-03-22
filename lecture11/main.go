package main

import "fmt"

type Book struct{
	Title string 
	Author string 
	ISBN string
	Price float32
	Pages int
}

func main(){
	var b1 Book 
	b1.Title ="Go Book"
	b1.Author ="Caleb Droxy"
	b1.ISBN = "12233-347"
	b1.Price = 150.34
	b1.Pages = 300

	fmt.Println(b1)
//fmt.Println("Hello Go")
}