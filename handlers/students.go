package handlers

import (
	"net/http"
)

func ListStudents(w http.ResponseWriter, r *http.Request) {
	s, err := studentRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_index.html", s)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	s, err := studentRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_create.html", s)
}

func InsertStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		err := studentRepository.Insert(name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/students", http.StatusSeeOther)
	}
}

func EditStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	s, err := studentRepository.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_edit.html", s)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		err := studentRepository.Update(id, name)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/students", http.StatusSeeOther)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := studentRepository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)
}
