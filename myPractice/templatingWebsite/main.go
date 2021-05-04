package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/docs", docs)
	http.HandleFunc("/features", features)
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/res/", http.StripPrefix("/res/", fs))

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `Welcome here in the loop of limitless`)
}

func docs(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `Welcome here in the loop of limitless`)
}
func features(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `Welcome here in the loop of limitless`)
}

// func docs(w http.ResponseWriter, r *http.Request) {
// 	ptmp, err := template.ParseFiles("template/base.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	//ptmp.Execute(w, nil)
// 	ptmp, err = template.ParseFiles("wpage/docs.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	ptmp.Execute(w, nil)
// 	//fmt.Fprintf(w, `Welcome here in the loop of limitless`)
// }

// func features(w http.ResponseWriter, r *http.Request) {
// 	ptmp, err := template.ParseFiles("template/base.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	//ptmp.Execute(w, nil)

// 	ptmp, err = template.ParseFiles("wpage/docs.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	ptmp.Execute(w, nil)
// 	//fmt.Fprintf(w, `Welcome here in the loop of limitless`)
// }
