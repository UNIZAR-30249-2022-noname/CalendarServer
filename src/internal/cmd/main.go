package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/docs"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//SetupRouter is a func which bind each uri with a handler function
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	horariosrv := horariosrv.New(nil)
	horarioHandler := handlers.NewHTTPHandler(horariosrv)
	r.GET("/ping", handlers.Ping)
	r.GET("/availableHours", horarioHandler.GetAvailableHours)
	r.POST("/newEntry", horarioHandler.PostNewEntry)

	return r
}

func main() {
	// · Swagger ·
	docs.SwaggerInfo.Title = "API UNIZAR calendar and schedule"
	docs.SwaggerInfo.Description = "This is API for managing and visulizating the calendar and schedule of Unizar."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080/swagger/index.html"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := SetupRouter()

	r.Run(":8080")
}
