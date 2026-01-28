package models

// StudentCourse represents a many-to-many relation between students and courses.
type StudentCourse struct {
	StudentID int
	CourseID  int
}

// StudentCourseView represents a student–course relation with
// human-readable fields, typically populated via JOIN queries.
type StudentCourseView struct {
	StudentID   int
	CourseID    int
	StudentName string
	CourseTitle string
}

// StudentsCoursesData is a view model for students–courses pages.
type StudentsCoursesData struct {
	Students []Student
	Courses  []Course
}
