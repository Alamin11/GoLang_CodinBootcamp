package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}
