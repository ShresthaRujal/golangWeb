package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy")
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d) // forward slach for routing child url
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)

	// or
	// http.Handle("/dog/", d)
	// http.Handle("/cat", c)
	// http.ListenAndServe(":8080", nil)

	//or
	// http.HandleFunc("/dog/", e)
	// http.HandleFunc("/cat", f)
	// http.ListenAndServe(":8080", nil)

	// or
	// http.Handle("/dog/", http.HandlerFunc(e))
	// http.Handle("/cat", http.HandlerFunc(f))
	// http.ListenAndServe(":8080", nil)

	//or using third party library
	//mux := httprouter.New()
	//mux.GET("/",index)
	//mux.GET("/blog/:catalog/:article",blogread)
}
func e(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy")
}

func f(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty")
}
