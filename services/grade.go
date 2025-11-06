package services

import (
	"course-enrollment-system/config"
	"database/sql"
)

func AssignGrade(enrollmentId int, grade string, remarks string) (int, error) {
	var id int

	query := `
		INSERT INTO grades
		(enrollment_id, grade, remarks)
		VALUES($1, $2, $3)
		RETURNING id
	`

	err := config.DB.QueryRow(query, enrollmentId, grade, remarks).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GradeEnrollmentIdExist(enrollmentId int) (bool, error) {
	isExist := true
	var id int

	query := `
		SELECT id
		FROM grades
		WHERE enrollment_id = $1
	`

	if err := config.DB.QueryRow(query, enrollmentId).Scan(&id); err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return true, err
	}

	return isExist, nil
}
