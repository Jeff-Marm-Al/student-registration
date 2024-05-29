package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/students", createStudent)
	server.GET("/students", getAllStudents)
	server.GET("/students/student/:firstname/:lastname", getStudentInfo)
	server.PUT("/students/student/:firstname/:lastname", updateStudentInfo)
	server.POST("/courses", createCourse)
	server.GET("/courses", getAllCourses)
}
