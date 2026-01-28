package repositories

import (
	"database/sql"
	"my-app-go/models"
)

// StudentCourseRepository defines database operations for working with
// a table of students and their courses.
type StudentCourseRepository interface {
	// GetAll returns all the student-course entities.
	GetAll() ([]models.StudentCourseView, error)

	// GetStudentsAndCourses returns StudentsCoursesData object that contains
	// all students and courses.
	GetStudentsAndCourses() (models.StudentsCoursesData, error)

	// Insert creates a new student-course entity with the given student ID
	// and course ID.
	Insert(studentID, courseID string) error

	// Delete removes a student-course entity by student ID
	// and course ID.
	Delete(studentID, courseID string) error
}

// studentCourseRepo is a SQL-based implementation of StudentCourseRepository.
type studentCourseRepo struct {
	DB *sql.DB
}

// NewStudentCourseRepository creates a new StudentCourseRepository backed
// by an SQL database.
func NewStudentCourseRepository(db *sql.DB) StudentCourseRepository {
	return &studentCourseRepo{DB: db}
}

func (r *studentCourseRepo) GetAll() ([]models.StudentCourseView, error) {
	rows, err := r.DB.Query(`
		SELECT sc.student_id, sc.course_id, s.name, c.title
		FROM students_courses sc
		JOIN students s ON sc.student_id = s.id
		JOIN courses c ON sc.course_id = c.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scList []models.StudentCourseView
	for rows.Next() {
		var sc models.StudentCourseView
		if err := rows.Scan(&sc.StudentID, &sc.CourseID, &sc.StudentName,
			&sc.CourseTitle); err != nil {

			return nil, err
		}
		scList = append(scList, sc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return scList, nil
}

func (r *studentCourseRepo) GetStudentsAndCourses() (
	models.StudentsCoursesData, error) {

	students, err := r.DB.Query("SELECT id, name FROM students")
	data := models.StudentsCoursesData{}
	if err != nil {
		return data, err
	}
	defer students.Close()

	courses, err := r.DB.Query("SELECT id, title FROM courses")
	if err != nil {
		return data, err
	}
	defer courses.Close()

	for students.Next() {
		var s models.Student
		if err := students.Scan(&s.ID, &s.Name); err != nil {
			return models.StudentsCoursesData{}, err
		}
		data.Students = append(data.Students, s)
	}

	if err := students.Err(); err != nil {
		return models.StudentsCoursesData{}, err
	}

	for courses.Next() {
		var c models.Course
		if err := courses.Scan(&c.ID, &c.Title); err != nil {
			return models.StudentsCoursesData{}, err
		}
		data.Courses = append(data.Courses, c)
	}

	if err := courses.Err(); err != nil {
		return models.StudentsCoursesData{}, err
	}

	return data, nil
}

func (r *studentCourseRepo) Insert(studentID, courseID string) error {
	_, err := r.DB.Exec(`INSERT INTO students_courses (student_id, course_id) 
			VALUES (?, ?)`, studentID, courseID)

	return err
}

func (r *studentCourseRepo) Delete(studentID, courseID string) error {
	_, err := r.DB.Exec(`DELETE FROM students_courses WHERE student_id = ? 
		AND course_id = ?`, studentID, courseID)

	return err
}
