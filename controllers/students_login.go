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

func StudentsLogin(c *gin.Context) {
	var payload dto.LoginPayload
	var response dto.StudentLoginResponse

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, errUser := services.GetUsersByUsernameWithPassword(payload.Username)

	if errUser == sql.ErrNoRows {
		utils.Error(c, http.StatusBadRequest, constants.ErrorMessage003)
		return
	} else if errUser != nil {
		utils.Error(c, http.StatusInternalServerError, errUser.Error())
		return
	}

	if passwordValid := utils.ComparePassword(payload.Password, user.Password); !passwordValid {
		utils.Error(c, http.StatusUnauthorized, constants.ErrorMessage004)
		return
	}

	studentId, errStudentId := services.GetStudentId(user.ID)

	if errStudentId != nil {
		utils.Error(c, http.StatusInternalServerError, errStudentId.Error())
		return
	}

	token, errToken := utils.GenerateJWT(payload.Username, studentId)

	if errToken != nil {
		utils.Error(c, http.StatusInternalServerError, errToken.Error())
		return
	}

	response.Username = payload.Username
	response.Token = token
	response.StudentId = studentId

	utils.Success(c, "Login Success", response)
}
