package models

type StudentCourse struct {
	StudentID int
	CourseID  int
}

type StudentCourseView struct {
	StudentID   int
	CourseID    int
	StudentName string
	CourseTitle string
}
