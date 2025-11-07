package services

import (
	"course-enrollment-system/config"
	"course-enrollment-system/dto"
	"course-enrollment-system/entity"
)

func CreateCourse(title, description string, teacherId int) (int, error) {
	var id int
	queryRow := `
		INSERT INTO courses
		(title, description, teacher_id)
		VALUES($1, $2, $3)
		RETURNING id
	`

	if err := config.DB.QueryRow(queryRow, title, description, teacherId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func GetAllCourses() ([]entity.Course, error) {
	var data []entity.Course
	query := `
		SELECT id, title, description, teacher_id
		FROM courses
	`

	courses, err := config.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for courses.Next() {
		var course entity.Course

		errCourse := courses.Scan(&course.ID, &course.Title, &course.Description, &course.TeacherID)

		if errCourse != nil {
			return nil, errCourse
		}

		data = append(data, course)
	}

	return data, nil
}

func GetDetailCourses(id string) (dto.CourseDetail, error) {
	var data dto.CourseDetail

	queryRow := `
		SELECT 
			c.id, 
			c.title,
			c.description,
			c.teacher_id,
			t.name AS teacher_name,
			c.created_at
		FROM courses AS c
		JOIN teachers AS t
		ON c.teacher_id = t.id
		WHERE c.id=$1
	`

	err := config.DB.QueryRow(queryRow, id).Scan(&data.Id, &data.Title, &data.Descriptions, &data.TeacherId, &data.TeacherName, &data.CreatedAt)

	if err != nil {
		return data, err
	}

	return data, nil
}

func GetCoursesByTeacherId(teacherId int) ([]entity.Course, error) {
	var data []entity.Course

	query := `
		SELECT id, title, description, teacher_id
		FROM courses
		WHERE teacher_id=$1
	`

	courses, err := config.DB.Query(query, teacherId)

	if err != nil {
		return data, err
	}

	for courses.Next() {
		var course entity.Course
		if errScan := courses.Scan(&course.ID, &course.Title, &course.Description, &course.TeacherID); errScan != nil {
			return data, errScan
		}

		data = append(data, course)
	}

	return data, nil
}

func GetCoursesByStudentId(studentId int) ([]entity.Course, error) {
	var data []entity.Course

	query := `
		SELECT
			c.id,
			c.title,
			c.description,
			c.teacher_id
		FROM enrollments AS e
		JOIN courses AS c ON e.course_id = c.id
		JOIN students AS s ON e.student_id = s.id
		WHERE student_id=$1
	`

	rows, err := config.DB.Query(query, studentId)

	if err != nil {
		return data, err
	}

	for rows.Next() {
		var row entity.Course

		errScan := rows.Scan(&row.ID, &row.Title, &row.Description, &row.TeacherID)

		if errScan != nil {
			return data, errScan
		}

		data = append(data, row)
	}

	return data, nil
}

func GetCourseDetailByStudentId(studentId int, courseId string) (entity.CoursesListByStudent, error) {
	var data entity.CoursesListByStudent

	query := `
		SELECT 
			c.id,
			c.title,
			c.description,
			c.teacher_id,
			t.name AS teacher_name,
			g.grade,
			g.remarks,
			e.id AS enrollment_id
		FROM enrollments AS e
		JOIN courses AS c ON e.course_id = c.id
		JOIN teachers AS t ON c.teacher_id = t.id
		LEFT JOIN grades AS g ON e.id = g.enrollment_id
		WHERE e.student_id=$1 AND c.id=$2
	`

	err := config.DB.QueryRow(query, studentId, courseId).Scan(&data.Id, &data.Title, &data.Descriptions, &data.TeacherId, &data.TeacherName, &data.Grade, &data.Remarks, &data.EnrollmentId)

	if err != nil {
		return data, err
	}

	return data, nil
}
