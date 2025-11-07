package middleware

import (
	"course-enrollment-system/constants"
	"course-enrollment-system/services"
	"course-enrollment-system/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func TeacherValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.Error(c, http.StatusUnauthorized, "missing or invalid token")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(tokenString)

		if err != nil {
			utils.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			utils.Error(c, http.StatusUnauthorized, "expired token")
			c.Abort()
			return
		}

		teacher, errTeacher := services.GetUsersByUsername(claims.Username)

		if errTeacher != nil {
			utils.Error(c, http.StatusInternalServerError, errTeacher.Error())
			c.Abort()
			return
		} else if teacher.Role != constants.Teacher {
			utils.Error(c, http.StatusUnauthorized, constants.ErrorMessage005)
			c.Abort()
			return
		}

		c.Set(constants.Username, claims.Username)
		c.Set(constants.Id, claims.MyId)
		c.Next()
	}
}
