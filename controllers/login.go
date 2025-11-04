package controllers

import (
	"database/sql"
	"net/http"
	"vehicle-service-api/dto"
	"vehicle-service-api/services"
	"vehicle-service-api/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var payload dto.LoginPayload
	var response dto.LoginResponse

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, errUser := services.GetUsersByUsernameWithPassword(payload.Username)

	if errUser == sql.ErrNoRows {
		utils.Error(c, http.StatusBadRequest, "Invalid username")
		return
	} else if errUser != nil {
		utils.Error(c, http.StatusInternalServerError, errUser.Error())
		return
	}

	if passwordValid := utils.ComparePassword(payload.Password, user.Password); !passwordValid {
		utils.Error(c, http.StatusUnauthorized, "Invalid password")
		return
	}

	token, errToken := utils.GenerateJWT(payload.Username)

	if errToken != nil {
		utils.Error(c, http.StatusInternalServerError, errToken.Error())
		return
	}

	response.Username = payload.Username
	response.Token = token

	utils.Success(c, "Login Success", response)
}
