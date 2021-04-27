package main

import (
	"fmt"
	"os"
)

func main() {
	//Showing directory path of working diectory
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(dir)

	//Creating directory
	// posf, err := os.Create("newFile.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer posf.Close()

	// n, err := posf.Write([]byte("This a file that is created using go lang"))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(n, err)

	//Calling file creating function
	// content := "fileName and contet are taking from the argument"
	// isFalse := createFile("another2.txt", content)
	// fmt.Println(isFalse)

	//checking if file is existing or not
	// fi, err := os.Stat("another.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(fi.ModTime().Clock())
	// fmt.Println(fi.ModTime().Date())
	// fmt.Println(fi.ModTime().GobEncode())
	// fmt.Println(fi.IsDir())
	// fmt.Println(fi.ModTime())
	// fmt.Println(fi.Name())
	// fmt.Println(fi.Size())
	// fmt.Println(fi.Sys())

	//Creating a folder
	err := os.Mkdir("lec25", 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(dirName)
}

//Creating a function to create file
func createFile(fileName, fileContent string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer posf.Close()

	_, err = posf.Write([]byte(fileContent))
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(n, err)
	return true
}
