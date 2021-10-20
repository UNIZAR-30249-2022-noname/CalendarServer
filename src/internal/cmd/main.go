package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//horarioHandler := handlers.NewHTTPHandler(nil)
	r.GET("/ping", handlers.Ping)
	//r.GET("/availableHours", horarioHandler.GetAvailableHours)

	return r
}

func main() {

	r := SetupRouter()

	r.Run(":8080")
}
