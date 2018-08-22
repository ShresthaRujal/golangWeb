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

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Computer programming", "four"},
				course{"CSCI-41", "Introduction to C programming", "six"},
			},
		},

		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-40", "Introduction to Computer programming", "four"},
				course{"CSCI-41", "Introduction to C programming", "six"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
