package handlers

import (
	"database/sql"
	"my-app-go/models"
	"net/http"
)

func ListCourses(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT courses.id, courses.title, teachers.ID, teachers.name
		FROM courses
		LEFT JOIN teachers ON courses.teacher_id = teachers.id
	`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var c models.Course
		var tID sql.NullInt64
		var tName sql.NullString
		if err := rows.Scan(&c.ID, &c.Title, &tID, &tName); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if tID.Valid {
			c.TeacherID = int(tID.Int64)
		}
		if tName.Valid {
			c.TeacherName = tName.String
		}
		courses = append(courses, c)
	}
	tpl.ExecuteTemplate(w, "courses_index.html", courses)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
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
	tpl.ExecuteTemplate(w, "courses_create.html", teachers)
}

func InsertCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		teacherID := r.FormValue("teacher_id")
		_, err := db.Exec("INSERT INTO courses (title, teacher_id) VALUES (?, ?)", title, teacherID)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/courses", http.StatusSeeOther)
	}
}

func EditCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	row := db.QueryRow("SELECT id, title, teacher_id FROM courses WHERE id = ?", id)
	var c models.Course
	if err := row.Scan(&c.ID, &c.Title, &c.TeacherID); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

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

	tpl.ExecuteTemplate(w, "courses_edit.html", struct {
		Course   models.Course
		Teachers []models.Teacher
	}{c, teachers})
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		title := r.FormValue("title")
		teacherID := r.FormValue("teacher_id")
		_, err := db.Exec("UPDATE courses SET title = ?, teacher_id = ? WHERE id = ?", title, teacherID, id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/courses", http.StatusSeeOther)
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM courses WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/courses", http.StatusSeeOther)
}
