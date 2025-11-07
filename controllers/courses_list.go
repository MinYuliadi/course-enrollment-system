package controllers

import (
	"course-enrollment-system/constants"
	"course-enrollment-system/dto"
	"course-enrollment-system/entity"
	"course-enrollment-system/services"
	"course-enrollment-system/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CoursesList(c *gin.Context) {
	var response []dto.CoursesList

	courseList, errList := services.GetAllCourses()

	if errList != nil {
		utils.Error(c, http.StatusInternalServerError, errList.Error())
		return
	}

	for _, course := range courseList {
		var temp dto.CoursesList

		temp.Id = course.ID
		temp.Title = course.Title
		temp.Descriptions = *course.Description
		temp.TeacherId = *course.TeacherID

		response = append(response, temp)
	}

	utils.Success(c, "Success", response)
}

func CourseListByTeacherId(c *gin.Context) {
	var response []dto.CoursesList
	teacherId, exist := c.Get(constants.Id)

	if !exist {
		utils.Error(c, http.StatusInternalServerError, constants.ErrorMessage006)
		return
	}

	courses, errCourses := services.GetCoursesByTeacherId(teacherId.(int))

	if errCourses == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, constants.ErrorMessage008)
		return
	}

	for _, course := range courses {
		var temp dto.CoursesList

		temp.Id = course.ID
		temp.Title = course.Title
		temp.Descriptions = course.Title
		temp.TeacherId = *course.TeacherID

		response = append(response, temp)
	}

	utils.Success(c, "Success", response)
}

func CourseListByStudentId(c *gin.Context) {
	var response []dto.CoursesList
	studentId, exist := c.Get(constants.Id)

	if !exist {
		utils.Error(c, http.StatusInternalServerError, constants.ErrorMessage006)
		return
	}

	courses, errCourses := services.GetCoursesByStudentId(studentId.(int))

	if errCourses != nil {
		utils.Error(c, http.StatusInternalServerError, errCourses.Error())
		return
	}

	for _, course := range courses {
		var temp dto.CoursesList

		temp.Id = course.ID
		temp.Title = course.Title
		temp.Descriptions = *course.Description
		temp.TeacherId = *course.TeacherID

		response = append(response, temp)
	}

	utils.Success(c, "Success", response)
}

func CourseDetailByStudentId(c *gin.Context) {
	var response dto.CoursesListByStudentDTO
	studentId, exist := c.Get(constants.Id)
	courseId := c.Param("id")
	if !exist {
		utils.Error(c, http.StatusInternalServerError, constants.ErrorMessage006)
		return
	}

	course, errCourse := services.GetCourseDetailByStudentId(studentId.(int), courseId)

	if errCourse == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, constants.ErrorMessage008)
		return
	}
	if errCourse != nil {
		utils.Error(c, http.StatusInternalServerError, errCourse.Error())
		return
	}

	response.Id = course.Id
	response.Title = course.Title
	response.Descriptions = course.Descriptions
	response.TeacherId = course.TeacherId
	response.TeacherName = course.TeacherName
	response.Grade = course.Grade
	response.Remarks = course.Remarks

	attendance, errAttendance := services.GetAttendanceByEnrollmentId(course.EnrollmentId)

	if errAttendance != nil {
		utils.Error(c, http.StatusInternalServerError, errAttendance.Error())
		return
	}

	response.Attendance = &[]entity.AttendanceListByEnrollment{}

	for _, attend := range attendance {
		var attendanceTemp entity.AttendanceListByEnrollment
		attendanceTemp.ID = attend.ID
		attendanceTemp.Date = attend.Date
		attendanceTemp.Status = attend.Status
		*response.Attendance = append(*response.Attendance, attendanceTemp)
	}

	utils.Success(c, "Success", response)
}
