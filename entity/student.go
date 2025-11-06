package entity

import "time"

type Student struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UserID    int       `db:"user_id" json:"user_id"`
}

type StudentsListByCourseEntity struct {
	Id           int    `db:"id" binding:"required"`
	Name         string `db:"name"`
	Email        string `db:"email"`
	CreatedAt    string `db:"created_at"`
	UserId       int    `db:"user_id"`
	EnrollmentId *int   `db:"enrollment_id"`
}
