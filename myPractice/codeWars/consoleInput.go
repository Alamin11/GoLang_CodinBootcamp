package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type whatever you want : ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You have typed : %q\n", input)
	fmt.Print("Now insert your birth year : ")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years of old at 2021", 2021-age)

}
