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

// CreateStudent displays a student creation form
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	s, err := studentRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_create.html", s)
}

// InsertStudent saves a new student
func InsertStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	err := studentRepository.Insert(name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)
}

// EditStudent displays a student editing form
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
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	err := studentRepository.Update(id, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)
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
