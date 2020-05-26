package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("something.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(welcome))
	http.Handle("/avenger", http.HandlerFunc(foo))
	http.Handle("/profile", http.HandlerFunc(bar))
	http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	/*tp, err := template.ParseFiles("something.gohtml")

	if err != nil{
		 log.Fatal("crashed")
	}*/

	err := tpl.ExecuteTemplate(w,"something.gohtml", "avengers, assemble in profile")

	if err != nil {
		log.Fatalln(err)
	}
}
