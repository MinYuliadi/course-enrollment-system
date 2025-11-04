package controllers

import (
	"database/sql"
	"net/http"
	"vehicle-service-api/dto"
	"vehicle-service-api/services"
	"vehicle-service-api/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var payload dto.RegisterPayload
	var responseData dto.RegisterResponse

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, errUser := services.GetUsersByUsername(payload.Username)

	if user.Username == payload.Username {
		utils.Error(c, http.StatusBadRequest, "username already exist")
		return
	}

	if errUser != nil && errUser != sql.ErrNoRows {
		utils.Error(c, http.StatusInternalServerError, errUser.Error())
		return
	}

	userId, errCreateUser := services.CreateUser(payload.Username, payload.Password)

	if errCreateUser != nil {
		utils.Error(c, http.StatusInternalServerError, errCreateUser.Error())
		utils.Logging(errCreateUser.Error())
		return
	}

	responseData.RegisteredId = userId

	utils.Success(c, "Registration successful", responseData)
}
