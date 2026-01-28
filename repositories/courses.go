package repositories

import (
	"database/sql"
	"my-app-go/models"
)

// CourseRepository defines database operations for working with courses.
type CourseRepository interface {
	// GetAll returns all courses with teacher information.
	GetAll() ([]models.Course, error)

	// Insert creates a new course with the given title and teacher ID.
	Insert(title, teacherID string) error

	// GetCoursesAndTeachers returns a course by ID
	// along with all teachers (used for edit forms).
	GetCourseAndTeachers(id string) (models.Course, []models.Teacher, error)

	// Update updates an existing course by ID.
	Update(id, title, teacherID string) error

	// Delete removes a course by ID.
	Delete(id string) error
}

// courseRepo is a SQL-based implementation of CourseRepository.
type courseRepo struct {
	DB *sql.DB
}

// NewCourseRepository creates a new CourseRepository backed by an SQL database.
func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepo{DB: db}
}

func (r *courseRepo) GetAll() ([]models.Course, error) {
	rows, err := r.DB.Query(`
		SELECT c.id, c.title, t.id, t.name
		FROM courses c
		LEFT JOIN teachers t ON c.teacher_id = t.id
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *courseRepo) Insert(title, teacherID string) error {
	_, err := r.DB.Exec(`
		INSERT INTO courses (title, teacher_id) 
		VALUES (?, ?)`, title, teacherID)

	return err
}

func (r *courseRepo) GetCourseAndTeachers(id string) (models.Course,
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

	if err := rows.Err(); err != nil {
		return c, nil, err
	}

	return c, teachers, nil
}

func (r *courseRepo) Update(id, title, teacherID string) error {
	_, err := r.DB.Exec("UPDATE courses SET title = ?, teacher_id = ? WHERE id = ?",
		title, teacherID, id)

	return err
}

func (r *courseRepo) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM courses WHERE id = ?", id)

	return err
}
