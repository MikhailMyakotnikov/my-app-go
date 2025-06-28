package handlers

import (
	"net/http"
)

// @Summary	Отображает список преподавателей
// @Description загружает HTML-страницу со списком преподавателей
// @Tags teachers
// @Produce text/html
// @Success 200 {string} string "HTML-страница"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /teachers [get]
func ListTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := teacherRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "teachers_index.html", teachers)
}

// @Summary Отображает форму создания преподавателя
// @Description Загружает HTML-страницу с формой для создания нового преподавателя
// @Tags teachers
// @Produce text/html
// @Success 200 {string} string "HTML-форма"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /teachers/create [get]
func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	teachers, err := teacherRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "teachers_create.html", teachers)
}

// @Summary Создает нового преподавателя
// @Description Получает данные из формы и сохраняет нового преподавателя в базу данных
// @Tags teachers
// @Accept application/x-www-form-urlencoded
// @Produce text/html
// @Param name formData string true "Имя преподавателя"
// @Success 303 {string} string "Редирект на /teachers"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /teachers/insert [post]
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

// @Summary Отображает форму редактирования преподавателя
// @Description Загружает HTML-страницу с формой для редактирования информации о преподавателе
// @Tags teachers
// @Produce text/html
// @Param id query string true "ID преподавателя"
// @Success 200 {string} string "HTML-форма редактирования"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /teachers/edit [get]
func EditTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	t, err := teacherRepository.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.ExecuteTemplate(w, "teachers_edit.html", t)
}

// @Summary Обновляет информацию о преподавателе
// @Description Обновляет имя преподавателя по ID, переданному через форму
// @Tags teachers
// @Accept application/x-www-form-urlencoded
// @Produce text/html
// @Param id formData string true "ID преподавателя"
// @Param name formData string true "Новое имя преподавателя"
// @Success 303 {string} string "Редирект на /teachers"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /teachers/update [post]
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

// @Summary Удаляет информацию о преподавателе
// @Description Удаляет имя преподавателя по ID, переданному через query-параметр
// @Tags teachers
// @Produce text/html
// @Param id query string true "ID преподавателя"
// @Success 303 {string} string "Редирект на /teachers"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /teachers/delete [get]
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := teacherRepository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/teachers", http.StatusSeeOther)
}
