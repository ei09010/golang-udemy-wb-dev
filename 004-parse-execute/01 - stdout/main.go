package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatal("Error while parsing file")
	}

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatal("error when passing file content to stdout", err)
	}

}
