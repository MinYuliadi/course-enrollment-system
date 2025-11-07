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

func TeachersRegister(c *gin.Context) {
	var payload dto.RegisterPayload
	var responseData dto.RegisterResponse

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, errUser := services.GetUsersByUsername(payload.Username)

	if user.Username == payload.Username {
		utils.Error(c, http.StatusBadRequest, constants.ErrorMessage001)
		return
	}

	if errUser != nil && errUser != sql.ErrNoRows {
		utils.Error(c, http.StatusInternalServerError, errUser.Error())
		return
	}

	isEmailAvailable, errEmail := services.CheckTeacherEmail(payload.Email)

	if errEmail != nil && errEmail != sql.ErrNoRows {
		utils.Error(c, http.StatusInternalServerError, errEmail.Error())
		return
	} else if !isEmailAvailable {
		utils.Error(c, http.StatusBadRequest, constants.ErrorMessage002)
		return
	}

	userId, errCreateUser := services.CreateUser(payload.Username, payload.Password, constants.Teacher)

	if errCreateUser != nil {
		utils.Error(c, http.StatusInternalServerError, errCreateUser.Error())
		return
	}

	teacherId, errCreateTeacher := services.CreateTeacher(payload.Name, payload.Email, userId)

	if errCreateTeacher != nil {
		utils.Error(c, http.StatusInternalServerError, errCreateTeacher.Error())
		return
	}

	responseData.RegisteredId = teacherId
	responseData.RegisteredUserId = userId

	utils.Success(c, "Registration successful", responseData)
}
