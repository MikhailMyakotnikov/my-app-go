package handlers

import (
	"net/http"
)

func ListTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := teacherRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "teachers_index.html", teachers)
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	teachers, err := teacherRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "teachers_create.html", teachers)
}

func InsertTeacher(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		err := teacherRepository.Insert(name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/teachers", http.StatusSeeOther)
	}
}

func EditTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	t, err := teacherRepository.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "teachers_edit.html", t)
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		err := teacherRepository.Update(id, name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/teachers", http.StatusSeeOther)
	}
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := teacherRepository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/teachers", http.StatusSeeOther)
}
