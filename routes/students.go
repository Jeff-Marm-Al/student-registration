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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create new student account.", "error": err.Error()})
		return 
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Student account created", "student": student})

}

func getAllStudents(context *gin.Context) {
	// calling the function to list all students
	students, err := models.GetAllStudents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch all students. Try again later.", "error": err.Error()})
		return 
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully retrieved all students", "students": students})
}

func getStudentInfo(context *gin.Context) {
	student, err := models.GetStudentInfo(context.Param("firstname"), context.Param("lastname"))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get student info. Try again later.", "error": err.Error()})
		return 
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully retrieved student info", "student": student})
}

func updateStudentInfo(context *gin.Context) {
	// gets the students info from the database
	student, err := models.GetStudentInfo(context.Param("firstname"), context.Param("lastname"))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get student info. Try again later.", "error": err.Error()})
		return 
	}

	// empty struct to update student info with
	var updatedStudent models.Student
	
	err = context.ShouldBindJSON(&updatedStudent) 	// payload of the api call applied to the empty struct

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse provided data", "error": err.Error()})
		return 
	}

	updatedStudent.StudentID = student.StudentID 	// assigns student ID to the updated student struct to keep the same

	err = updatedStudent.UpdateStudentInfo() // api call to actually update the student info

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update student info. Try again later.", "error": err.Error()})
		return 
	}
	
	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated student info", "student": updatedStudent})
}