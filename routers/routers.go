package routers

import (
	"vehicle-service-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	auth := api.Group("/auth")

	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	return router
}
