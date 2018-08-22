package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func logError(err error) {
	log.Fatalln(err)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	logError(err)
}
