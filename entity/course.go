package entity

import "time"

type Course struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description *string   `db:"description" json:"description"`
	TeacherID   *int      `db:"teacher_id" json:"teacher_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
