package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	// tpl, err := template.ParseFiles("tpl.gohtml")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// err = tpl.Execute(os.Stdout, nil)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	//Using execute func from parseGlob
	err := tpl.ExecuteTemplate(os.Stdout, "template1.gohtml", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = tpl.ExecuteTemplate(os.Stdout, "template2.gohtml", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = tpl.ExecuteTemplate(os.Stdout, "template3.gohtml", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
