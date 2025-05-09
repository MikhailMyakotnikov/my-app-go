package repositories

import (
	"database/sql"
	"my-app-go/models"
)

type StudentCourseRepository interface {
	GetAll() ([]models.StudentCourseView, error)
	GetStudentsAndCourses() (
		struct {
			Students []models.Student
			Courses  []models.Course
		}, error)
	Insert(student_id string, course_id string) error
	Delete(studentID string, courseID string) error
}

type studentCourseRepo struct {
	DB *sql.DB
}

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

	return scList, nil
}

func (r *studentCourseRepo) GetStudentsAndCourses() (struct {
	Students []models.Student
	Courses  []models.Course
}, error) {

	students, err := r.DB.Query("SELECT id, name FROM students")
	data := struct {
		Students []models.Student
		Courses  []models.Course
	}{}
	if err != nil {
		return data, err
	}
	defer students.Close()

	courses, err := r.DB.Query("SELECT id, title FROM courses")
	if err != nil {
		return data, err
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

	data = struct {
		Students []models.Student
		Courses  []models.Course
	}{
		Students: studentList,
		Courses:  courseList,
	}

	return data, nil
}

func (r *studentCourseRepo) Insert(studentID string, courseID string) error {
	_, err := r.DB.Exec(`INSERT INTO students_courses (student_id, course_id) 
			VALUES (?, ?)`, studentID, courseID)

	return err
}

func (r *studentCourseRepo) Delete(studentID string, courseID string) error {
	_, err := r.DB.Exec(`DELETE FROM students_courses WHERE student_id = ? 
		AND course_id = ?`, studentID, courseID)

	return err
}
