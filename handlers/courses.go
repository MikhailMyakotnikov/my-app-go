package handlers

import (
	"my-app-go/models"
	"net/http"
)

func ListCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := courseRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "courses_index.html", courses)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	teachers, err := teacherRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "courses_create.html", teachers)
}

func InsertCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		teacherID := r.FormValue("teacher_id")
		err := courseRepository.Insert(title, teacherID)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/courses", http.StatusSeeOther)
	}
}

func EditCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	c, teachers, err := courseRepository.GetCoursesAndTeachers(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
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
		err := courseRepository.Update(id, title, teacherID)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/courses", http.StatusSeeOther)
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := courseRepository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/courses", http.StatusSeeOther)
}
