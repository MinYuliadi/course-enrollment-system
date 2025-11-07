package controllers

import (
	"course-enrollment-system/constants"
	"course-enrollment-system/dto"
	"course-enrollment-system/services"
	"course-enrollment-system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TeacherCreateCourse(c *gin.Context) {
	var payload dto.CoursePayload
	var response dto.CourseResponse

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	teacherId, isExist := c.Get(constants.Id)

	if !isExist {
		utils.Error(c, http.StatusInternalServerError, constants.ErrorMessage006)
		return
	}

	courseId, errCourse := services.CreateCourse(payload.Title, payload.Description, teacherId.(int))

	if errCourse != nil {
		utils.Error(c, http.StatusInternalServerError, errCourse.Error())
		return
	}

	response.CourseId = courseId

	utils.Success(c, "Success create course", response)
}
