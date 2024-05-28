package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/students", createStudent)
	server.GET("/students", getAllStudents)
	server.GET("/students/:firstname/:lastname", getStudentInfo)
	server.POST("/courses", createCourse)
	server.GET("/courses", getAllCourses)
}
