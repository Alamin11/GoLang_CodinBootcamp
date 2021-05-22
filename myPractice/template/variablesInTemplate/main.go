package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Remove-confution ; Focus-what you like`)
	if err != nil {
		fmt.Println(err.Error())
	}
	// nf, err := os.Create("index.html")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer nf.Close()
	// io.Copy(nf, strings.NewReader())
}
