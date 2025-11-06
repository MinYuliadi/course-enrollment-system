package routers

import (
	"course-enrollment-system/controllers"
	"course-enrollment-system/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	students := api.Group("/students", middleware.StudentValidation())
	teachers := api.Group("/teachers", middleware.TeacherValidation())

	{
		api.GET("/courses", controllers.CoursesList)
		api.GET("/courses/:id", controllers.CourseDetail)

		api.POST("/register/students", controllers.StudentsRegister)
		api.POST("/login/students", controllers.StudentsLogin)
		api.POST("/register/teachers", controllers.TeachersRegister)
		api.POST("/login/teachers", controllers.TeachersLogin)

		students.POST("/enrollment", controllers.StudentsEnrollment)
		students.GET("/my-course", controllers.CourseListByStudentId)

		teachers.POST("/courses", controllers.TeacherCreateCourse)
		teachers.GET("/my-course", controllers.CourseListByTeacherId)
		teachers.GET("/courses/:id/students", controllers.StudentsListByCourseId)
		teachers.POST("/attendance", controllers.AssignAttendance)
	}

	return router
}
