package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/resourse/", http.StripPrefix("/resourse/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8888", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website")
}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website. This is my contact page")
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//fmt.Fprintf(w, "Welcome to my website")
	fmt.Fprintf(w, "<img src=\"resourse/venom.jpg\"width=\"750px\" height=\"500px\"/>")
}
