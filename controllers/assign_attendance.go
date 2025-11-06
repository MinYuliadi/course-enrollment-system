package controllers

import (
	"course-enrollment-system/constants"
	"course-enrollment-system/dto"
	"course-enrollment-system/services"
	"course-enrollment-system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssignAttendance(c *gin.Context) {
	var request dto.AssignAttendanceRequestDTO
	var response dto.AssignAttendanceResponseDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		return
	}

	if request.Status != "" && request.Status != "present" && request.Status != "absent" && request.Status != "late" {
		utils.Error(c, http.StatusBadRequest, constants.ErrorMessage010)
		return
	}

	attendanceId, err := services.AssignAttendance(request.EnrollmentId, request.Status, request.Date)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.AttendanceId = attendanceId
	utils.Success(c, "Attendance assigned successfully", response)
}
