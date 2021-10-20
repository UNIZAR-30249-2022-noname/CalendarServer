package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	horariosrv := horariosrv.New(nil)
	horarioHandler := handlers.NewHTTPHandler(horariosrv)
	r.GET("/ping", handlers.Ping)
	r.GET("/availableHours", horarioHandler.GetAvailableHours)

	return r
}

func main() {

	r := SetupRouter()

	r.Run(":8080")
}
