package services

import (
	"course-enrollment-system/config"
	"course-enrollment-system/entity"
)

func CreateStudents(name, email string, userId int) (int, error) {
	var id int
	query := `
		INSERT INTO students
		(name, email, user_id)
		VALUES($1, $2, $3)
		RETURNING id
	`

	if err := config.DB.QueryRow(query, name, email, userId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func GetStudentId(userId int) (int, error) {
	var id int
	query := `
		SELECT id
		FROM students
		WHERE user_id=$1
	`

	if err := config.DB.QueryRow(query, userId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func GetStudentsByCourseId(courseId string) ([]entity.StudentsListByCourseEntity, error) {
	var students []entity.StudentsListByCourseEntity

	query := `
		SELECT s.id, s.name, s.email, s.created_at, s.user_id, ec.id AS enrollment_id
		FROM enrollments ec
		LEFT JOIN students s ON ec.student_id = s.id
		WHERE ec.course_id = $1
	`

	rows, err := config.DB.Query(query, courseId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var student entity.StudentsListByCourseEntity
		if err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.CreatedAt, &student.UserId, &student.EnrollmentId); err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}
