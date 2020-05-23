package main

import (
	"text/template"
	"log"
	"os"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatal("Error while parsing file")
	}

	nf, err := os.Create("myFile.html")

	if err != nil {
		log.Fatal("Didn't manage to create file vital for output")
	}

	defer nf.Close()


	err = tpl.Execute(nf, nil)

	if err != nil {
		log.Fatal("error when passing file content to stdout", err)
	}
}
