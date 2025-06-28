package handlers

import (
	"my-app-go/models"
	"net/http"
)

// @Summary	Отображает список курсов
// @Description загружает HTML-страницу со списком курсов
// @Tags courses
// @Produce text/html
// @Success 200 {string} string "HTML-страница"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /courses [get]
func ListCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := courseRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "courses_index.html", courses)
}

// @Summary Отображает форму создания курса
// @Description Загружает HTML-страницу с формой для создания нового курса
// @Tags courses
// @Produce text/html
// @Success 200 {string} string "HTML-форма"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /courses/create [get]
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	teachers, err := teacherRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "courses_create.html", teachers)
}

// @Summary Создает новый курс
// @Description Получает данные из формы и сохраняет новый курс в базу данных
// @Tags courses
// @Accept application/x-www-form-urlencoded
// @Produce text/html
// @Param title formData string true "Название курса"
// @Param teacher_id formData string true "ID преподавателя"
// @Success 303 {string} string "Редирект на /courses"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /courses/insert [post]
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

// @Summary Отображает форму редактирования курса
// @Description Загружает HTML-страницу с формой для редактирования информации о курсе
// @Tags courses
// @Produce text/html
// @Param id query string true "ID курса"
// @Success 200 {string} string "HTML-форма редактирования"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /courses/edit [get]
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

// @Summary Обновляет информацию о курсе
// @Description Обновляет название курса по ID, переданному через форму
// @Tags courses
// @Accept application/x-www-form-urlencoded
// @Produce text/html
// @Param id formData string true "ID курса"
// @Param title formData string true "Новое название курса"
// @Param teacher_id formData string true "ID преподавателя"
// @Success 303 {string} string "Редирект на /courses"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /courses/update [post]
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

// @Summary Удаляет информацию о курсе
// @Description Удаляет название курса по ID, переданному через query-параметр
// @Tags courses
// @Produce text/html
// @Param id query string true "ID курса"
// @Success 303 {string} string "Редирект на /courses"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /courses/delete [get]
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := courseRepository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/courses", http.StatusSeeOther)
}
