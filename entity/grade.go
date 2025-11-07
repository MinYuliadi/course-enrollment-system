package entity

type Grade struct {
	ID           int     `db:"id" json:"id"`
	EnrollmentID int     `db:"enrollment_id" json:"enrollment_id"`
	Grade        *string `db:"grade" json:"grade"`
	Remarks      *string `db:"remarks" json:"remarks"`
}
