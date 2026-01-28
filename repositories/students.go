package repositories

import (
	"database/sql"
	"my-app-go/models"
)

// StudentRepository defines database operations for working with
// a table of students.
type StudentRepository interface {
	GetAll() ([]models.Student, error)

	// Insert creates a new student with the given name.
	Insert(name string) error

	// GetById gets information about a student by ID.
	GetById(id string) (models.Student, error)

	// Update updates an existing student by ID.
	Update(id, name string) error

	// Delete removes a student by ID.
	Delete(id string) error
}

// studentRepo is a SQL-based implementation of StudentRepository.
type studentRepo struct {
	DB *sql.DB
}

// NewStudentRepository creates a new StudentRepository backed by an SQL database.
func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentRepo{DB: db}
}

func (r *studentRepo) GetAll() ([]models.Student, error) {
	rows, err := r.DB.Query("SELECT id, name FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *studentRepo) Insert(name string) error {
	_, err := r.DB.Exec("INSERT INTO students (name) VALUES (?)", name)

	return err
}

func (r *studentRepo) GetById(id string) (models.Student, error) {
	row := r.DB.QueryRow("SELECT id, name FROM students WHERE id = ?", id)
	var s models.Student
	if err := row.Scan(&s.ID, &s.Name); err != nil {
		return s, err
	}

	return s, nil
}

func (r *studentRepo) Update(id, name string) error {
	_, err := r.DB.Exec("UPDATE students SET name = ? WHERE id = ?", name, id)

	return err
}

func (r *studentRepo) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM students WHERE id = ?", id)

	return err
}
