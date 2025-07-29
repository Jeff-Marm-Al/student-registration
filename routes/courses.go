package routes

import (
	"net/http"
	"github.com/Jeff-Marm-Al/student-registration/models"
	"github.com/gin-gonic/gin"
)

func createCourse(context *gin.Context) {
	var course models.Course

	err := context.ShouldBindJSON(&course)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse provided data", "error": err.Error()})
		return 
	}

	err = course.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create course.", "error": err.Error()})
		return 
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Course was created."})
}

func getAllCourses(context *gin.Context) {
	courses, err := models.GetAllCourses()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch list of courses", "error": err.Error()})
		return 
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully retrieved all courses", "courses": courses})
}