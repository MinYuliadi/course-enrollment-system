package dto

import "course-enrollment-system/entity"

type StudentsListByCourseDTO struct {
	Id           int                                  `json:"id" binding:"required"`
	Name         string                               `json:"name"`
	Email        string                               `json:"email"`
	CreatedAt    string                               `json:"created_at"`
	UserId       int                                  `json:"user_id"`
	EnrollmentId *int                                 `json:"enrollment_id"`
	Attendances  *[]entity.AttendanceListByEnrollment `json:"attendances"`
}
