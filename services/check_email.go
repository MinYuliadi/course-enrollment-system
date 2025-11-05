package services

import (
	"course-enrollment-system/config"
	"database/sql"
)

func CheckStudentEmail(email string) (bool, error) {
	var dbmail string
	query := `
		SELECT email
		FROM students
		WHERE email=$1
	`

	if err := config.DB.QueryRow(query, email).Scan(&dbmail); err == sql.ErrNoRows {
		return true, err
	} else if err != nil {
		return false, err
	}

	return false, nil
}

func CheckTeacherEmail(email string) (bool, error) {
	var dbmail string
	query := `
		SELECT email
		FROM teachers
		WHERE email=$1
	`

	if err := config.DB.QueryRow(query, email).Scan(&dbmail); err == sql.ErrNoRows {
		return true, err
	} else if err != nil {
		return false, err
	}

	return false, nil
}
