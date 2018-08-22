package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	rujal := sage{
		Name:  "Rujal",
		Motto: "Die or Die Doing",
	}

	err := tpl.Execute(os.Stdout, rujal)
	if err != nil {
		log.Fatalln(err)
	}
}
