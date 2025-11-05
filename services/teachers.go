package services

import "course-enrollment-system/config"

func CreateTeacher(name, email string, userId int) (int, error) {
	var id int
	query := `
		INSERT INTO teachers
		name=$1, email=$2, user_id=$3
		RETURNING id
	`

	if err := config.DB.QueryRow(query, name, email, userId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func GetTeacherId(userId int) (int, error) {
	var id int
	query := `
		SELECT id
		FROM teachers
		WHERE user_id=$1
	`

	if err := config.DB.QueryRow(query, userId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
