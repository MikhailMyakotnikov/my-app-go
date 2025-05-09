package handlers

import (
	"database/sql"
	"html/template"
	"my-app-go/repositories"
)

var (
	db                      *sql.DB
	tpl                     *template.Template
	courseRepository        repositories.CourseRepository
	teacherRepository       repositories.TeacherRepository
	studentRepository       repositories.StudentRepository
	studentCourseRepository repositories.StudentCourseRepository
)

// var db *sql.DB
// var tpl *template.Template
// var courseRepository repositories.CourseRepository = repositories.NewCourseRepository(db)
// var teacherRepository repositories.TeacherRepository = repositories.NewTeacherRepository(db)
// var studentRepository repositories.StudentRepository = repositories.NewStudentRepository(db)
// var studentCourseRepository repositories.StudentCourseRepository = repositories.NewStudentCourseRepository(db)

func Init(database *sql.DB, templates *template.Template) {
	db = database
	tpl = templates

	courseRepository = repositories.NewCourseRepository(db)
	teacherRepository = repositories.NewTeacherRepository(db)
	studentRepository = repositories.NewStudentRepository(db)
	studentCourseRepository = repositories.NewStudentCourseRepository(db)
}
