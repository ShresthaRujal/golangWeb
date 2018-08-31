package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	dbTest, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mysql")
	check(err)
	db = dbTest
	defer db.Close()

	err = db.Ping()
	check(err)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi, Welcome to Rujal's Sql Practice")
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Created Table customer ", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES("James");`)
	check(err)

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Inserted Into Table customer", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM customer")
	check(err)

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)

		fmt.Println(name)
		fmt.Fprintln(w, "Retrived record :", name)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="rujal";`)
	check(err)

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Deleted record ", n)
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name= "rujal" WHERE name="James";`)
	check(err)

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Updated record ", n)
}
