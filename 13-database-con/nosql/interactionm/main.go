package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type user struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	ID     bson.ObjectId `json:"id" bson:"_id"`
}

var db *mgo.Database
var se *mgo.Session

func main() {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	se = s
	r := httprouter.New()
	r.GET("/", index)
	// http.GEt("/user/", insert)
	r.GET("/read/:id", read)
	r.GET("/test/:id", test)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "this is index page")
}

func test(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "test "+p.ByName("id"))
	if !bson.IsObjectIdHex(p.ByName("id")) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}
}

// func insert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	u := user{
// 		Name:   "Rujal",
// 		Gender: "Male",
// 	}
// 	//for getting user from form
// 	// json.NewDecoder(r.Body).Decode(&u)

// 	//create id
// 	u.ID = bson.NewObjectId()

// 	//getseesion, insert u in users in database mongodb
// 	se.DB("mongodb").C("users").Insert(u)

// 	uj, _ := json.Marshal(u)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, uj)
// }

func read(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// id := "5b878f9e88c36b1a7cfca269"
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}
	oid := bson.ObjectIdHex(id)
	u := user{}
	err := se.DB("mongodb").C("users").Find(oid).One(&u)

	if err != nil {
		fmt.Println(err)
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
