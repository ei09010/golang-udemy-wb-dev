package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct{
	Name string
	Age int
	Powers []string
}
func init() {
	tpl = template. Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := person{"thor", 500, []string{"flying", "thunder attack"}}

	myOutputFile, err:= os.Create("profile.html")
	if err != nil{
		log.Fatal("error creating output file")
	}
	defer myOutputFile.Close()

	err = tpl.ExecuteTemplate(myOutputFile, "tpl.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
