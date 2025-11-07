package dto

type AssignGradeRequestDTO struct {
	EnrollmentId int    `json:"enrollment_id" binding:"required"`
	Grade        string `json:"grade" binding:"required"`
	Remarks      string `json:"remarks"`
}

type AssignGradeResponseDTO struct {
	GradeId int `json:"grade_id"`
}
