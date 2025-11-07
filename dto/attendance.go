package dto

type AssignAttendanceRequestDTO struct {
	EnrollmentId int    `json:"enrollment_id" binding:"required"`
	Status       string `json:"status" binding:"required"`
	Date         string `json:"date" binding:"required"`
}

type AssignAttendanceResponseDTO struct {
	AttendanceId int `json:"attendance_id"`
}
