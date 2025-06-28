package handlers

import (
	"net/http"
)

// @Summary	Отображает список студентов
// @Description загружает HTML-страницу со списком студентов
// @Tags students
// @Produce text/html
// @Success 200 {string} string "HTML-страница"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students [get]
func ListStudents(w http.ResponseWriter, r *http.Request) {
	s, err := studentRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_index.html", s)
}

// @Summary Отображает форму создания студента
// @Description Загружает HTML-страницу с формой для создания нового студента
// @Tags students
// @Produce text/html
// @Success 200 {string} string "HTML-форма"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students/create [get]
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	s, err := studentRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_create.html", s)
}

// @Summary Создает нового студента
// @Description Получает данные из формы и сохраняет нового студента в базу данных
// @Tags students
// @Accept application/x-www-form-urlencoded
// @Produce text/html
// @Param name formData string true "Имя студента"
// @Success 303 {string} string "Редирект на /students"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students/insert [post]
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

// @Summary Отображает форму редактирования студента
// @Description Загружает HTML-страницу с формой для редактирования информации о студенте
// @Tags students
// @Produce text/html
// @Param id query string true "ID студента"
// @Success 200 {string} string "HTML-форма редактирования"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students/edit [get]
func EditStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	s, err := studentRepository.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "students_edit.html", s)
}

// @Summary Обновляет информацию о студенте
// @Description Обновляет имя студента по ID, переданному через форму
// @Tags students
// @Accept application/x-www-form-urlencoded
// @Produce text/html
// @Param id formData string true "ID студента"
// @Param name formData string true "Новое имя студента"
// @Success 303 {string} string "Редирект на /students"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students/update [post]
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

// @Summary Удаляет информацию о студенте
// @Description Удаляет имя студента по ID, переданному через query-параметр
// @Tags students
// @Produce text/html
// @Param id query string true "ID студента"
// @Success 303 {string} string "Редирект на /students"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /students/delete [get]
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := studentRepository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)
}
