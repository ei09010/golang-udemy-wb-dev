package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", chien)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func foo(w http.ResponseWriter, req *http.Request){

	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("dog.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w,"dog.gohtml",nil)
}

func chien(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "dog.jpg")
}
