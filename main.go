package main

import (
	"github.com/Jeff-Marm-Al/student-registration/db"
	"github.com/Jeff-Marm-Al/student-registration/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // this is localhost:8080
}
