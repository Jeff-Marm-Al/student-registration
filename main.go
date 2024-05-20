package main

import (
	"fmt"

	"github.com/Jeff-Marm-Al/student-registration/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Run(":8080") // this is localhost:8080
}