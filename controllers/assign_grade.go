package controllers

import (
	"course-enrollment-system/constants"
	"course-enrollment-system/dto"
	"course-enrollment-system/services"
	"course-enrollment-system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssignGrade(c *gin.Context) {
	var request dto.AssignGradeRequestDTO
	var response dto.AssignGradeResponseDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	isExist, err := services.GradeEnrollmentIdExist(request.EnrollmentId)

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if isExist {
		utils.Error(c, http.StatusBadRequest, constants.ErrorMessage010)
		return
	}

	gradeId, err := services.AssignGrade(request.EnrollmentId, request.Grade, request.Remarks)

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.GradeId = gradeId

	utils.Success(c, "Success", response)
}
