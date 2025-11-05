package dto

type EnrollmentPayload struct {
	CourseId int `json:"courseId"`
}

type EnrollmentResponse struct {
	EnrollmentId int `json:"enrollmentId"`
}
