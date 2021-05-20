package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//name := "Mohammad Al amin"
	name := os.Args[1] //storing the data taken from user screen
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	tmp := `
	<!DOCTYPE html>
	<html lang="eng">
		<head>
		<title>First template</title>
		</head>
		<body>
		<h1>Hello ` + name + `</h1>
		</body>
	</html>
	`
	//creating a file and taking data from user using os.Create() package
	nf, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tmp))
	fmt.Println(tmp)
}
