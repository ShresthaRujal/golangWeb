package main

import (
	"net/http"
	"text/template"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		sid := uuid.Must(uuid.NewV4())
		c = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		http.SetCookie(w, c)
	}

	//if user already exits
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	//get cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.html", u)
}
