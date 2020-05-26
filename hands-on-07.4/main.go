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

	http.HandleFunc("/", dogList)
	// resources is indicated in the image path from index
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func dogList(w http.ResponseWriter, request *http.Request) {

	err := tpl.Execute(w, nil)

	if err != nil{

		log.Fatal(err)
	}
}
