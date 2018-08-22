package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

var fm = template.FuncMap{
	"uc":   strings.ToUpper,
	"ft":   firstThree,
	"cv":   intToString,
	"sqrt": sqRoot,
}

func sqRoot(i int) float64 {
	return math.Pow(float64(i), 2)
}

func intToString(i int) string {
	str := strconv.Itoa(i)
	return str
}
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var tpl *template.Template

func init() {

	tpl = template.Must(template.New(`test.gohtml`).Funcs(fm).ParseFiles("test.gohtml"))
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
		Speed:       2600,
	}

	b := car{
		Manufacture: "Ferrari",
		Speed:       3600,
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
