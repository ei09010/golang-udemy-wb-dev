package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}


func main() {

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", randomTemplateFunc)
	http.ListenAndServe(":8080", nil)
}


func randomTemplateFunc(w http.ResponseWriter, _ *http.Request){

	err := tpl.ExecuteTemplate(w,"index.gohtml", nil)

	if err != nil{
		log.Fatal(err)
	}
}
