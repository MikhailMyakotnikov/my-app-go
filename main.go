package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"my-app-go/handlers"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var tpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		// Example: user:123@tcp(127.0.0.1:1234)/my_app_go
		"%s:%s@tcp(%s:%s)/%s?tls=skip-verify",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	handlers.Init(db, tpl)

	// index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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

	fmt.Println("Server started at http://0.0.0.0:8081")
	err = http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Fatal("Server startup error: ", err)
	}
}
