package main

import (
	"io"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database
var se *mgo.Session

func main() {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	se = s
	http.HandleFunc("/", index)
	http.HandleFunc("/user/", getData)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is index page")
}

func getData(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is getUser")
}

func insert(w http.ResponseWriter, r *http.Request){
	db.
}