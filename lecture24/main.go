package main

import (
	"fmt"
	"os"
)

func main() {
	//Showing directory path of working diectory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)

	//Creating directory
	posf, err := os.Create("newFile.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer posf.Close()

	n, err := posf.Write([]byte("This a file that is created using go lang"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(n, err)
	createFile("another.txt")

}

//Creating a function to create file
func createFile(fileName string) {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer posf.Close()

	n, err := posf.Write([]byte("Creating a file from the function"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(n, err)
}
