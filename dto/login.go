package dto

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type StudentLoginResponse struct {
	Username  string `json:"username"`
	StudentId int    `json:"studentId"`
	Token     string `json:"token"`
}

type TeacherLoginResponse struct {
	Username  string `json:"username"`
	TeacherId int    `json:"teacherId"`
	Token     string `json:"token"`
}
