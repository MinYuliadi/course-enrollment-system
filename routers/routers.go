package routers

import (
	"course-enrollment-system/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	students := api.Group("/students")
	teachers := api.Group("/teachers")

	{
		students.POST("/register", controllers.StudentsRegister)
		students.POST("/login", controllers.StudentsLogin)

		teachers.POST("/register", controllers.TeachersRegister)
		teachers.POST("/login", controllers.TeachersLogin)
	}

	return router
}
