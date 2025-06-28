package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"my-app-go/handlers"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "my-app-go/docs"

	_ "github.com/go-sql-driver/mysql"
)

//		@title			My CRUD App API
//		@version		1.0
//		@description	API для управления курсами и студентами
//	 	@license.name	MIT

var db *sql.DB
var tpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	var err error
	db, err = sql.Open("mysql", "mikhail:123qwe@tcp(127.0.0.1:3306)/my_app_go?tls=custom")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	handlers.Init(db, tpl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/teachers", http.StatusSeeOther)
	})

	// teachers
	http.HandleFunc("/teachers", handlers.ListTeachers)
	http.HandleFunc("/teachers/create", handlers.CreateTeacher)
	http.HandleFunc("/teachers/insert", handlers.InsertTeacher)
	http.HandleFunc("/teachers/edit", handlers.EditTeacher)
	http.HandleFunc("/teachers/update", handlers.UpdateTeacher)
	http.HandleFunc("/teachers/delete", handlers.DeleteTeacher)

	// courses
	http.HandleFunc("/courses", handlers.ListCourses)
	http.HandleFunc("/courses/create", handlers.CreateCourse)
	http.HandleFunc("/courses/insert", handlers.InsertCourse)
	http.HandleFunc("/courses/edit", handlers.EditCourse)
	http.HandleFunc("/courses/update", handlers.UpdateCourse)
	http.HandleFunc("/courses/delete", handlers.DeleteCourse)

	// students
	http.HandleFunc("/students", handlers.ListStudents)
	http.HandleFunc("/students/create", handlers.CreateStudent)
	http.HandleFunc("/students/insert", handlers.InsertStudent)
	http.HandleFunc("/students/edit", handlers.EditStudent)
	http.HandleFunc("/students/update", handlers.UpdateStudent)
	http.HandleFunc("/students/delete", handlers.DeleteStudent)

	// students_courses
	http.HandleFunc("/students_courses", handlers.ListStudentsCourses)
	http.HandleFunc("/students_courses/create", handlers.CreateStudentCourse)
	http.HandleFunc("/students_courses/insert", handlers.InsertStudentCourse)
	http.HandleFunc("/students_courses/delete", handlers.DeleteStudentCourse)

	fmt.Println("Server started at http://localhost:8080")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("Server startup error: ", err)
	}

	http.Handle("/swagger/", httpSwagger.WrapHandler)
}
