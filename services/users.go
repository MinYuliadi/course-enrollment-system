package services

import (
	"course-enrollment-system/config"
	"course-enrollment-system/entity"
	"course-enrollment-system/utils"
)

func GetUsersByUsername(username string) (user entity.User, err error) {
	query := `
		SELECT id, username, role
		FROM users
		WHERE username=$1
	`

	if err = config.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Role); err != nil {
		return user, err
	}

	return user, nil
}

func GetUsersByUsernameWithPassword(username string) (user entity.User, err error) {
	query := `
		SELECT id, username, role, password
		FROM users
		WHERE username=$1
	`

	if err = config.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Role, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(username, password, role string) (id int, err error) {
	query := `
		INSERT INTO users
		(username, password, role)
		VALUES($1, $2, $3)
		RETURNING id
	`

	hashedPassword, errHashing := utils.HashPassword(password)

	if errHashing != nil {
		return 0, errHashing
	}

	if errQuery := config.DB.QueryRow(query, username, hashedPassword, role).Scan(&id); errQuery != nil {
		return 0, errQuery
	}

	return id, nil
}
