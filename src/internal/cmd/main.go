package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/docs"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
