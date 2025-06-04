package handlers

import (
	"net/http"
)

func ListStudentsCourses(w http.ResponseWriter, r *http.Request) {
	scList, err := studentCourseRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_courses_index.html", scList)
}

func CreateStudentCourse(w http.ResponseWriter, r *http.Request) {
	data, err := studentCourseRepository.GetStudentsAndCourses()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_courses_create.html", data)
}

func InsertStudentCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		studentID := r.FormValue("student_id")
		courseID := r.FormValue("course_id")

		err := studentCourseRepository.Insert(studentID, courseID)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/students_courses", http.StatusSeeOther)
	}
}

func DeleteStudentCourse(w http.ResponseWriter, r *http.Request) {
	studentID := r.URL.Query().Get("student_id")
	courseID := r.URL.Query().Get("course_id")

	err := studentCourseRepository.Delete(studentID, courseID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/students_courses", http.StatusSeeOther)
}
