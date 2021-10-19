package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", handlers.Ping)

	return r
}

func main() {

	r := setupRouter()

	r.Run(":8080")
}
