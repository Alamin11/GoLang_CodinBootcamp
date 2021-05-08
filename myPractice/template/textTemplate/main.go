// package main

// import (
// 	"fmt"
// 	"os"
// 	"text/template"
// )

// func main() {
// 	msg := `Hello {{.Name}} , We welcome you here .
// 	You can see more if you want , go {{.WebLink}}

// 	This is my first text template .
// 	Welcome to my world of love .
// 	You have successfully signed up . Veryfy {{.VLink}}.
// 	Here you will find the actual meaning of life.

// 	regards,
// 	Md. Al amin
// 	`

// 	value := struct {
// 		Name    string
// 		WebLink string
// 		VLink   string
// 	}{
// 		"Al amin",
// 		"http://http.myportfolio.com",
// 		"http://knoweMe.com",
// 	}

// 	ptml := template.New("email")

// 	eml, err := ptml.Parse(msg)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	eml.Execute(os.Stdout, value)

// }

package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	msg := `
	Hello {{.FirstName}} {{.LastName}},
	Welcome to my website.
	You need to verify first {{.VLink}}
	You can now explore on {{.WebLink}}

	regards
	{{.firstName}}
	CEO at my website
	`

	val := struct {
		FirstName string
		LastName  string
		VLink     string
		WebLink   string
	}{
		"Mohammad",
		"Al amin",
		"http://verify.com",
		"http://mywebsite.com",
	}

	ptml := template.New("greet")

	ptm, err := ptml.Parse(msg)
	if err != nil {
		fmt.Println(err.Error())
	}

	ptm.Execute(os.Stdout, val)
}
