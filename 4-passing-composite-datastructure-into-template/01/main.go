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
func main() {
	// sages := []string{"Rujal", "Shrestha", "Madhyapur", "Thimi"}
	//map
	sages := map[string]string{
		"name":     "Rujal",
		"lastname": "Shrestha",
		"Address":  "Madhyapur Thimi",
		"Phone":    "606060",
	}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
