package handlers

import (
	"my-app-go/models"
	"net/http"
)

func ListTeachers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM teachers")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var t models.Teacher
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		teachers = append(teachers, t)
	}
	tpl.ExecuteTemplate(w, "teachers_index.html", teachers)
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "teachers_create.html", nil)
}

func InsertTeacher(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		_, err := db.Exec("INSERT INTO teachers (name) VALUES (?)", name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/teachers", http.StatusSeeOther)
	}
}

func EditTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := db.QueryRow("SELECT id, name FROM teachers WHERE id = ?", id)

	var t models.Teacher
	if err := row.Scan(&t.ID, &t.Name); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "teachers_edit.html", t)
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		_, err := db.Exec("UPDATE teachers SET name = ? WHERE id = ?", name, id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/teachers", http.StatusSeeOther)
	}
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/teachers", http.StatusSeeOther)
}
