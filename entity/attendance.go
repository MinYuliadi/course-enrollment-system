package entity

import "time"

type Attendance struct {
	ID           int       `db:"id" json:"id"`
	EnrollmentID int       `db:"enrollment_id" json:"enrollment_id"`
	Date         time.Time `db:"date" json:"date"`
	Status       string    `db:"status" json:"status"`
}
