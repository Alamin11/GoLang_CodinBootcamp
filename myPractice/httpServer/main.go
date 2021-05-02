package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8888", nil)

}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my website , Happy exploring `)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my website , Happy exploring about me `)
}
