package main

import (
	"fmt"
	"net/http"
)

func main() {
	//var x string
	//handler func(ResponseWriter, *Request)
	//handle func() eta amader network e joto req ashe sob handle kore

	//request handler
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/others", others)
	//The following line's code will start Go’s default HTTP server
	//and listen for connections on port 8080.
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
