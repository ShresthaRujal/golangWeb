package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type hotdog int

var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//exploring request
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   url.Values
		Header        http.Header
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.ContentLength,
	}

	tpl.ExecuteTemplate(w, "test.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
