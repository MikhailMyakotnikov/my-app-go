package handlers

import (
	"my-app-go/models"
	"net/http"
)

func ListStudentsCourses(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT sc.student_id, sc.course_id, s.name, c.title
		FROM students_courses sc
		JOIN students s ON sc.student_id = s.id
		JOIN courses c ON sc.course_id = c.id
	`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var scList []models.StudentCourseView
	for rows.Next() {
		var sc models.StudentCourseView
		if err := rows.Scan(&sc.StudentID, &sc.CourseID, &sc.StudentName, &sc.CourseTitle); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		scList = append(scList, sc)
	}
	tpl.ExecuteTemplate(w, "students_courses_index.html", scList)
}

func CreateStudentCourse(w http.ResponseWriter, r *http.Request) {
	students, err := db.Query("SELECT id, name FROM students")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer students.Close()

	courses, err := db.Query("SELECT id, title FROM courses")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer courses.Close()

	var studentList []models.Student
	for students.Next() {
		var s models.Student
		students.Scan(&s.ID, &s.Name)
		studentList = append(studentList, s)
	}

	var courseList []models.Course
	for courses.Next() {
		var c models.Course
		courses.Scan(&c.ID, &c.Title)
		courseList = append(courseList, c)
	}

	data := struct {
		Students []models.Student
		Courses  []models.Course
	}{
		Students: studentList,
		Courses:  courseList,
	}

	tpl.ExecuteTemplate(w, "students_courses_create.html", data)
}

func InsertStudentCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		studentID := r.FormValue("student_id")
		courseID := r.FormValue("course_id")

		_, err := db.Exec("INSERT INTO students_courses (student_id, course_id) VALUES (?, ?)", studentID, courseID)
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

	_, err := db.Exec("DELETE FROM students_courses WHERE student_id = ? AND course_id = ?", studentID, courseID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/students_courses", http.StatusSeeOther)
}
