package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fmt.Println("\nfile: ", f, "\nheader: ", h, "\nerr", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		//store on a file
		err = os.Mkdir("user", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// io.WriteString(w, `<form method="post" enctype="multipart/form-data"> <input type="file" name="q"> <input type="submit"> </form> <br> `+s)
	tpl.ExecuteTemplate(w, "index.html", s)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
