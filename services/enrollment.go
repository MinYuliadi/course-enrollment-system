package services

import (
	"course-enrollment-system/config"
	"course-enrollment-system/entity"
	"database/sql"
)

func CreateEnrollment(studentId, courseId int) (int, error) {
	var id int
	query := `
		INSERT INTO enrollments
		(student_id, course_id)
		VALUES($1, $2)
		RETURNING id
	`

	err := config.DB.QueryRow(query, studentId, courseId).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetEnrollmentByStudentId(studentId int) ([]entity.Enrollment, error) {
	var data []entity.Enrollment

	query := `
		SELECT id, student_id, course_id
		FROM enrollments
		WHERE student_id=$1
	`

	enrollments, err := config.DB.Query(query, studentId)
	if err != nil && err != sql.ErrNoRows {
		return data, err
	}

	for enrollments.Next() {
		var enrollment entity.Enrollment
		if err := enrollments.Scan(&enrollment.ID, &enrollment.StudentID, &enrollment.CourseID); err != nil {
			return data, err
		}

		data = append(data, enrollment)
	}

	return data, nil
}
