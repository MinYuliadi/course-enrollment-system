package dto

import "time"

type CoursePayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CourseResponse struct {
	CourseId int `json:"courseId"`
}

type CoursesList struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"description"`
	TeacherId    int    `json:"teacher_id"`
}

type CourseDetail struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Descriptions string    `json:"description"`
	TeacherId    int       `json:"teacher_id"`
	TeacherName  string    `json:"teacher_name"`
	CreatedAt    time.Time `json:"created_at"`
}
