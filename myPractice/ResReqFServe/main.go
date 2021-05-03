package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/res/", http.StripPrefix("/res/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Well come to my website's home page ")
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//fmt.Fprintf(w, "<img src=\"venom.jpg\"/>")

	fmt.Fprintf(w, "<img src=\"res/venom.jpg\" width=\"700px\" height=\"500px\">")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Well come to my website's contact page ")
}
