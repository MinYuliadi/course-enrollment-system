package controllers

import (
	"course-enrollment-system/constants"
	"course-enrollment-system/dto"
	"course-enrollment-system/services"
	"course-enrollment-system/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentsEnrollment(c *gin.Context) {
	var payload dto.EnrollmentPayload
	var response dto.EnrollmentResponse

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	studentId, isExist := c.Get(constants.Id)

	if !isExist {
		utils.Error(c, http.StatusInternalServerError, constants.ErrorMessage006)
		return
	}

	enrollmentsById, errId := services.GetEnrollmentByStudentId(studentId.(int))

	if errId != sql.ErrNoRows && errId != nil {
		utils.Error(c, http.StatusInternalServerError, errId.Error())
		return
	}

	var isDuplicate = false

	for _, enrollment := range enrollmentsById {
		if enrollment.CourseID == payload.CourseId {
			isDuplicate = true
			break
		}
	}

	if isDuplicate {
		utils.Error(c, http.StatusBadRequest, constants.ErrorMessage009)
		return
	}

	enrollmentId, err := services.CreateEnrollment(studentId.(int), payload.CourseId)

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.EnrollmentId = enrollmentId

	utils.Success(c, "Enrollment Success", response)
}
