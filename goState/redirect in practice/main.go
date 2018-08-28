package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("your request method at foo: ", r.Method, "\n\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("your request at bar:", r.Method)
	//changes method from post to get (303)
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther)
	// or
	// http.Redirect(w, r, "/", http.StatusSeeOther)

	//dont change method (307)
	http.Redirect(w, r, "/", 307)
	//move permanently 301
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("your request at barred: ", r.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}
