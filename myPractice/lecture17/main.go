package main

import (
	"fmt"
	"net/http"
)

func main() {
	//var x string
	//handler func(ResponseWriter, *Request)

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/others", others)
	http.ListenAndServe(":8080", nil)

}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome To my first golang webpage`)

}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome To my first about page`)

}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome To my contact page`)

}
func others(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome To my others page`)

}
