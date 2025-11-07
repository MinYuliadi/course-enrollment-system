package controllers

import (
	"course-enrollment-system/dto"
	"course-enrollment-system/entity"
	"course-enrollment-system/services"
	"course-enrollment-system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentsListByCourseId(c *gin.Context) {
	var response []dto.StudentsListByCourseDTO
	id := c.Param("id")

	students, err := services.GetStudentsByCourseId(id)

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	for _, student := range students {
		var temp dto.StudentsListByCourseDTO

		temp.Id = student.Id
		temp.Name = student.Name
		temp.Email = student.Email
		temp.CreatedAt = student.CreatedAt
		temp.UserId = student.UserId
		temp.EnrollmentId = student.EnrollmentId
		temp.Grade = student.Grade
		temp.Remarks = student.Remarks

		attendaces, err := services.GetAttendanceByEnrollmentId(student.EnrollmentId)

		if err != nil {
			utils.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		temp.Attendances = &[]entity.AttendanceListByEnrollment{}

		for _, attendance := range attendaces {
			var attendanceTemp entity.AttendanceListByEnrollment
			attendanceTemp.ID = attendance.ID
			attendanceTemp.Date = attendance.Date
			attendanceTemp.Status = attendance.Status

			*temp.Attendances = append(*temp.Attendances, attendanceTemp)
		}

		response = append(response, temp)
	}

	utils.Success(c, "Success", response)
}
