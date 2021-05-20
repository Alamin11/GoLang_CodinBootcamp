package main

import "fmt"

func main() {
	name := "Mohammad Al amin"
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
	fmt.Println(tmp)
}
