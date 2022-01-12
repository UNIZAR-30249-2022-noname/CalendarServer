package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/docs"
	"github.com/gin-contrib/cors"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//SetupRouter is a func which bind each uri with a handler function
func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	horariorepo := horarioRepositorio.New()
	horariosrv := horariosrv.New(horariorepo)
	horarioHandler := handlers.NewHTTPHandler(horariosrv)
	r.GET("/ping", handlers.Ping)
	r.GET("/availableHours", horarioHandler.GetAvailableHours)
	r.POST("/updateScheduler", horarioHandler.PostUpdateScheduler)
	r.GET("/listDegrees", horarioHandler.ListDegrees)
	r.GET("/getEntries", horarioHandler.GetEntries)
	r.GET("/getICS", horarioHandler.GetICS)
	r.POST("/updateByCSV", horarioHandler.UpdateByCSV)

	return r
}

func main() {
	// · Swagger ·
	docs.SwaggerInfo.Title = "API UNIZAR calendar and schedule"
	docs.SwaggerInfo.Description = "This is API for managing and visulizating the calendar and schedule of Unizar."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := SetupRouter()

	r.Run(":8080")
}
