package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {

	case "/dog":
		io.WriteString(w, "doggy")

	case "/cat":
		io.WriteString(w, "kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
