package main

import "fmt"

//Defining a struct type
type Address struct {
	name       string
	street     string
	city       string
	state      string
	postalCode int
}

/*type Address struct{
	name, street, city, state string
	postalCode int
}*/

func main() {

	//Initialization of address type data and accessing using struct literal
	var a = Address{name: "Al amin", street: "Kobi Sukanto Shorok", city: "Jhenaidah", state: "khulna", postalCode: 7330}
	fmt.Println("Here is a struct 'a' :  ", a)
	//accessing a single element of struct using dot notation
	x := a.name
	fmt.Println("name : ", x)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a2 := Address{city: "Chuadanga"}
	fmt.Println("New address :", a2)

	//Annonymus struct using struct literals
	struct1 := struct {
		name       string
		age        int
		occupation string
		salary     float32
	}{"Alamin", 28, "Teacher", 30000.0}

	fmt.Println(struct1)

	//custom data type accessing from another file
	var v1 Ami
	v1 = 10.0
	fmt.Println(v1)

}
