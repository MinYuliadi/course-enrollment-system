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

func CourseDetail(c *gin.Context) {
	var response dto.CourseDetail
	id := c.Param("id")

	course, err := services.GetDetailCourses(id)

	if err == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, constants.ErrorMessage008)
		return
	} else if err != nil && err != sql.ErrNoRows {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response = course
	utils.Success(c, "Success", response)
}
