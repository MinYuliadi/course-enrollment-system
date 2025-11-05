package entity

import "time"

type Enrollment struct {
	ID         int       `db:"id" json:"id"`
	StudentID  int       `db:"student_id" json:"student_id"`
	CourseID   int       `db:"course_id" json:"course_id"`
	EnrolledAt time.Time `db:"enrolled_at" json:"enrolled_at"`
}
