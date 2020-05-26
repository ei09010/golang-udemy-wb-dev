package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	ck, err := req.Cookie("visitCounter")

	if err == http.ErrNoCookie {

		ck = &http.Cookie{
			Name:  "visitCounter",
			Value: "0",
		}
	}
	cookieValue, _ := strconv.Atoi(ck.Value)

	cookieValue++

	ck.Value = strconv.Itoa(cookieValue)

	http.SetCookie(w, ck)

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
