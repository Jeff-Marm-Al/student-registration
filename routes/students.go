package routes

import (
	"net/http"

	"github.com/Jeff-Marm-Al/student-registration/models"
	"github.com/gin-gonic/gin"
)

func createStudent(context *gin.Context) {
	// created var student struct
	var student models.Student

	err := context.ShouldBindJSON(&student) // payload of the api call is applied to the student struct variable

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data", "error": err.Error()})
		return 
	}

	err = student.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create new student account."})
		return 
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Student account created", "event": student})

}

func getAllStudents(context *gin.Context) {
	// calling the function to list all students
	students, err := models.GetAllStudents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch all students. Try again later."})
		return 
	}

	context.JSON(http.StatusOK, students)
}