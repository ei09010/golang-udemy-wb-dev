package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/avenger", foo)
	http.HandleFunc("/profile", bar)
	http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "avengers, assemble in profile")
}