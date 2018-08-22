package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml")) // must does errorchecking and return template
}

func logError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	logError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	logError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	logError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	logError(err)

	err = tpl.Execute(os.Stdout, nil)
	logError(err)
}
