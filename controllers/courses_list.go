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

func CoursesList(c *gin.Context) {
	var response []dto.CoursesList

	courseList, errList := services.GetAllCourses()

	if errList != nil {
		utils.Error(c, http.StatusInternalServerError, errList.Error())
		return
	}

	for _, course := range courseList {
		var temp dto.CoursesList

		temp.Id = course.ID
		temp.Title = course.Title
		temp.Descriptions = *course.Description
		temp.TeacherId = *course.TeacherID

		response = append(response, temp)
	}

	utils.Success(c, "Success", response)
}

func CourseListByTeacherId(c *gin.Context) {
	var response []dto.CoursesList
	teacherId, exist := c.Get(constants.Id)

	if !exist {
		utils.Error(c, http.StatusInternalServerError, constants.ErrorMessage006)
		return
	}

	courses, errCourses := services.GetCoursesByTeacherId(teacherId.(int))

	if errCourses == sql.ErrNoRows {
		utils.Error(c, http.StatusNotFound, constants.ErrorMessage008)
		return
	}

	for _, course := range courses {
		var temp dto.CoursesList

		temp.Id = course.ID
		temp.Title = course.Title
		temp.Descriptions = course.Title
		temp.TeacherId = *course.TeacherID

		response = append(response, temp)
	}

	utils.Success(c, "Success", response)
}

func CourseListByStudentId(c *gin.Context) {
	var response []dto.CoursesList
	studentId, exist := c.Get(constants.Id)

	if !exist {
		utils.Error(c, http.StatusInternalServerError, constants.ErrorMessage006)
		return
	}

	courses, errCourses := services.GetCoursesByStudentId(studentId.(int))

	if errCourses != nil {
		utils.Error(c, http.StatusInternalServerError, errCourses.Error())
		return
	}

	for _, course := range courses {
		var temp dto.CoursesList

		temp.Id = course.ID
		temp.Title = course.Title
		temp.Descriptions = *course.Description
		temp.TeacherId = *course.TeacherID

		response = append(response, temp)
	}

	utils.Success(c, "Success", response)
}
