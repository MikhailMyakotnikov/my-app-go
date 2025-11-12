package repositories

import (
	"database/sql"
	"my-app-go/models"
)

type TeacherRepository interface {
	GetAll() ([]models.Teacher, error)
	Insert(name string) error
	GetById(id string) (models.Teacher, error)
	Update(id string, name string) error
	Delete(id string) error
}

type teacherRepo struct {
	DB *sql.DB
}

func NewTeacherRepository(db *sql.DB) TeacherRepository {
	return &teacherRepo{DB: db}
}

func (r *teacherRepo) GetAll() ([]models.Teacher, error) {
	rows, err := r.DB.Query("SELECT id, name FROM teachers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var t models.Teacher
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		teachers = append(teachers, t)
	}

	return teachers, nil
}

func (r *teacherRepo) Insert(name string) error {
	_, err := r.DB.Exec("INSERT INTO teachers (name) VALUES (?)", name)

	return err
}

func (r *teacherRepo) GetById(id string) (models.Teacher, error) {
	row := r.DB.QueryRow("SELECT id, name FROM teachers WHERE id = ?", id)

	var t models.Teacher

	if err := row.Scan(&t.ID, &t.Name); err != nil {
		return t, err
	}

	return t, nil
}

func (r *teacherRepo) Update(id, name string) error {
	_, err := r.DB.Exec("UPDATE teachers SET name = ? WHERE id = ?", name, id)

	return err
}

func (r *teacherRepo) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM teachers WHERE id = ?", id)

	return err
}
