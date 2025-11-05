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

func TeachersLogin(c *gin.Context) {
	var payload dto.LoginPayload
	var response dto.TeacherLoginResponse

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

	teacherId, errTeacherId := services.GetTeacherId(user.ID)

	if errTeacherId != nil {
		utils.Error(c, http.StatusInternalServerError, errTeacherId.Error())
		return
	}

	token, errToken := utils.GenerateJWT(payload.Username, teacherId)

	if errToken != nil {
		utils.Error(c, http.StatusInternalServerError, errToken.Error())
		return
	}

	response.Username = payload.Username
	response.Token = token
	response.TeacherId = teacherId

	utils.Success(c, "Login Success", response)
}
