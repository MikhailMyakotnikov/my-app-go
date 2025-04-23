package handlers

import (
	"my-app-go/models"
	"net/http"
)

func ListStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM students")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		students = append(students, s)
	}
	tpl.ExecuteTemplate(w, "students_index.html", students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "students_create.html", nil)
}

func InsertStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		_, err := db.Exec("INSERT INTO students (name) VALUES (?)", name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/students", http.StatusSeeOther)
	}
}

func EditStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := db.QueryRow("SELECT id, name FROM students WHERE id = ?", id)

	var s models.Student
	if err := row.Scan(&s.ID, &s.Name); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_edit.html", s)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		_, err := db.Exec("UPDATE students SET name = ? WHERE id = ?", name, id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/students", http.StatusSeeOther)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)
}
