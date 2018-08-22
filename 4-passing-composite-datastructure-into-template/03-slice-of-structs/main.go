package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name    string
	Address string
}

type car struct {
	Manufacture string
	Speed       int
}

type item struct {
	Wisdom []sage
	Hope   []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	buddha := sage{
		Name:    "Buddha",
		Address: "Lumbini",
	}

	rujal := sage{
		Name:    "Rujal",
		Address: "Thimi",
	}

	sita := sage{
		Name:    "Sita",
		Address: "Janakpur",
	}

	a := car{
		Manufacture: "BMW",
		Speed:       260,
	}

	b := car{
		Manufacture: "Ferrari",
		Speed:       360,
	}

	sages := []sage{buddha, rujal, sita}
	cars := []car{a, b}

	// items := item{Wisdom: sages, Hope: cars}
	//or
	items := struct {
		Wisdom []sage
		Hope   []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, items)
	if err != nil {
		log.Fatalln(err)
	}
}
