package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/staticFile", content)
	//http.HandleFunc("/contact", contact)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/resourse/", http.StripPrefix("/resourse/", fs))

	//http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("assets"))))
	//fs := http.FileServer(http.Dir("assets/"))
	//http.Handle("/res/", http.StripPrefix("/res/", fs))
	http.ListenAndServe(":8888", nil)

}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my website , Happy exploring `)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my website , Happy exploring about me `)
}

func content(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "<img src=\"venom.jpg\">")

}

// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
// 	fmt.Fprintf(w, "<img src=\"res/venom.jpg\"/>")
// }
