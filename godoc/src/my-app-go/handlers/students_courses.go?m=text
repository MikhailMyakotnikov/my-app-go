package handlers

import (
	"net/http"
)

// @Summary	Отображает список студентов и их курсов
// @Description загружает HTML-страницу со списком студентов и курсов, которые они посещают
// @Tags students_courses
// @Produce text/html
// @Success 200 {string} string "HTML-страница"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students_courses [get]
func ListStudentsCourses(w http.ResponseWriter, r *http.Request) {
	scList, err := studentCourseRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_courses_index.html", scList)
}

// @Summary Отображает форму создания студентов и их курсов
// @Description Загружает HTML-страницу с формой для создания новой записи о студенте и курсе, который он посещает
// @Tags students_courses
// @Produce text/html
// @Success 200 {string} string "HTML-форма"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students_courses/create [get]
func CreateStudentCourse(w http.ResponseWriter, r *http.Request) {
	data, err := studentCourseRepository.GetStudentsAndCourses()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_courses_create.html", data)
}

// @Summary Создает новую запись о студенте и курсе, который он посещает
// @Description Получает данные из формы и сохраняет новую запись о студенте и курсе, который он посещает, в базу данных
// @Tags students_courses
// @Accept application/x-www-form-urlencoded
// @Produce text/html
// @Param student_id formData string true "ID студента"
// @Param course_id formData string true "ID курса"
// @Success 303 {string} string "Редирект на /students_courses"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students_courses/insert [post]
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

// @Summary Удаляет информацию о студенте и курсе, который он посещает
// @Description Удаляет информацию о студенте и курсе, который он посещает, по ID студента и ID курса, переданным через query-параметры
// @Tags students_courses
// @Produce text/html
// @Param student_id query string true "ID студента"
// @Param course_id query string true "ID курса"
// @Success 303 {string} string "Редирект на /students_courses"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students_courses/delete [get]
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
