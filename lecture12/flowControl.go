package main

import "fmt"

func main(){

//var age int
//fmt.Println("Enter your age : ")
//fmt.Scanf("%d", &age)

/*if age < 3{
	fmt.Println("Infant")
}else if age >= 3 && age < 13{
	fmt.Println("Children")
}else if age >= 13 && age <= 19{
	fmt.Println("Teenager")
}else{
	fmt.Println("Adult")
}*/

//Switch Case

/*switch age{

	case 2:
		fmt.Println("Infant")
	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("Children")
	fallthrough
	case 13, 14,15,16,17,18,19:
		fmt.Println("Teenage")
	default:
		fmt.Println("Adult")


}*/

//Loops
//Plain for loops 
/*for i := 0; i<10 ; i++ {
	fmt.Printf("Hello no %d\n", i)

}*/

//range loops 
 student := []int{1,2,3,4,5,6}
 for _, val := range student{
	fmt.Println(val)
}

//While like loop 
/*i:=1
for i<10 {
	fmt.Printf("Hello on %d \n", i)
	i++
}*/




}