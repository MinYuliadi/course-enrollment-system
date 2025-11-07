package services

import (
	"course-enrollment-system/config"
	"course-enrollment-system/entity"
	"time"
)

func AssignAttendance(enrollmentId int, status string, date string) (int, error) {
	var attendanceId int

	query := `
		INSERT INTO attendance
		(enrollment_id, status, date)
		VALUES($1, $2, $3)
		RETURNING id
	`

	parsedDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		return 0, err
	}

	if err := config.DB.QueryRow(query, enrollmentId, status, parsedDate).Scan(&attendanceId); err != nil {
		return 0, err
	}

	return attendanceId, nil
}

func GetAttendanceByEnrollmentId(enrollmentId int) ([]entity.AttendanceListByEnrollment, error) {
	var attendanceList []entity.AttendanceListByEnrollment
	query := `
		SELECT id, TO_CHAR(date, 'YYYY-MM-DD') AS date, status
		FROM attendance
		WHERE enrollment_id = $1
		ORDER BY date DESC
	`

	rows, err := config.DB.Query(query, enrollmentId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendance entity.AttendanceListByEnrollment
		if err := rows.Scan(&attendance.ID, &attendance.Date, &attendance.Status); err != nil {
			return nil, err
		}
		attendanceList = append(attendanceList, attendance)
	}

	return attendanceList, nil
}
