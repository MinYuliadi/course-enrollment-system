package entity

import "time"

type Enrollment struct {
	ID         int       `db:"id" json:"id"`
	StudentID  int       `db:"student_id" json:"student_id"`
	CourseID   int       `db:"course_id" json:"course_id"`
	EnrolledAt time.Time `db:"enrolled_at" json:"enrolled_at"`
}

type CoursesListByStudent struct {
	Id           int                           `db:"id"`
	Title        string                        `db:"title"`
	Descriptions string                        `db:"description"`
	TeacherId    int                           `db:"teacher_id"`
	TeacherName  string                        `db:"teacher_name"`
	Attendance   *[]AttendanceListByEnrollment `db:"attendances"`
	Grade        *string                       `db:"grade"`
	Remarks      *string                       `db:"remarks"`
	EnrollmentId int                           `db:"enrollment_id"`
}
