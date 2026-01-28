package models

type Course struct {
	ID        int
	Title     string
	TeacherID int

	// TeacherName is filled from joined teachers table.
	TeacherName string
}
