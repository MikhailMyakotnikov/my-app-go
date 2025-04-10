// main.go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID    int
	Name  string
	Email string
}

var db *sql.DB
var tpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	var err error
	db, err = sql.Open("mysql", "mikhail:123qwe@tcp(127.0.0.1:3306)/my_app_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM persons")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var persons []Person
	for rows.Next() {
		var p Person
		rows.Scan(&p.ID, &p.Name, &p.Email)
		persons = append(persons, p)
	}
	tpl.ExecuteTemplate(w, "index.html", persons)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create.html", nil)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		_, err := db.Exec("INSERT INTO persons (name, email) VALUES (?, ?)", name, email)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := db.QueryRow("SELECT * FROM persons WHERE id = ?", id)
	var p Person
	row.Scan(&p.ID, &p.Name, &p.Email)
	tpl.ExecuteTemplate(w, "edit.html", p)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")
		_, err := db.Exec("UPDATE persons SET name=?, email=? WHERE id=?", name, email, id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM persons WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
