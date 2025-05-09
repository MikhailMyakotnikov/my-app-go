package repositories

import (
	"database/sql"
	"my-app-go/models"
)

type CourseRepository interface {
	GetAll() ([]models.Course, error)
	Insert(title string, teacherID string) error
	GetCoursesAndTeachers(id string) (models.Course, []models.Teacher, error)
	Update(id string, title string, teacherID string) error
	Delete(id string) error
}

type courseRepo struct {
	DB *sql.DB
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepo{DB: db}
}

func (r *courseRepo) GetAll() ([]models.Course, error) {
	rows, err := r.DB.Query(`
		SELECT courses.id, courses.title, teachers.ID, teachers.name
		FROM courses
		LEFT JOIN teachers ON courses.teacher_id = teachers.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var c models.Course
		var tID sql.NullInt64
		var tName sql.NullString

		if err := rows.Scan(&c.ID, &c.Title, &tID, &tName); err != nil {
			return nil, err
		}
		if tID.Valid {
			c.TeacherID = int(tID.Int64)
		}
		if tName.Valid {
			c.TeacherName = tName.String
		}
		courses = append(courses, c)
	}

	return courses, nil
}

func (r *courseRepo) Insert(title string, teacherID string) error {
	_, err := r.DB.Exec(`
		INSERT INTO courses (title, teacher_id) 
		VALUES (?, ?)`, title, teacherID)

	return err
}

func (r *courseRepo) GetCoursesAndTeachers(id string) (models.Course,
	[]models.Teacher, error) {

	row := r.DB.QueryRow("SELECT id, title, teacher_id FROM courses WHERE id = ?", id)
	var c models.Course
	if err := row.Scan(&c.ID, &c.Title, &c.TeacherID); err != nil {
		return c, nil, err
	}

	rows, err := r.DB.Query("SELECT id, name FROM teachers")
	if err != nil {
		return c, nil, err
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var t models.Teacher
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return c, nil, err
		}
		teachers = append(teachers, t)
	}

	return c, teachers, nil
}

func (r *courseRepo) Update(id string, title string, teacherID string) error {
	_, err := r.DB.Exec("UPDATE courses SET title = ?, teacher_id = ? WHERE id = ?",
		title, teacherID, id)

	return err
}

func (r *courseRepo) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM courses WHERE id = ?", id)

	return err
}
