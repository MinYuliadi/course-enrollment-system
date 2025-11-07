package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
	Data      any    `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, statusCode int, message string, isSuccess bool, data any) {
	c.JSON(statusCode, Response{
		Code:      statusCode,
		Message:   message,
		IsSuccess: isSuccess,
		Data:      data,
	})

	Logging(message)
}

// Shortcut helpers for common responses
func Success(c *gin.Context, message string, data any) {
	SendResponse(c, http.StatusOK, message, true, data)
}

func Created(c *gin.Context, message string, data any) {
	SendResponse(c, http.StatusCreated, message, true, data)
}

func Error(c *gin.Context, statusCode int, message string) {
	SendResponse(c, statusCode, message, false, nil)
}
