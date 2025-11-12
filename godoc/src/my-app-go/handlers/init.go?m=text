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

func Init(database *sql.DB, templates *template.Template) {
	db = database
	tpl = templates

	courseRepository = repositories.NewCourseRepository(db)
	teacherRepository = repositories.NewTeacherRepository(db)
	studentRepository = repositories.NewStudentRepository(db)
	studentCourseRepository = repositories.NewStudentCourseRepository(db)
}
