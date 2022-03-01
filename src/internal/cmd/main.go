package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/docs"
	"github.com/gin-contrib/cors"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	uploaddata "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/uploadData"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	uploaddatarepositorymysql "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio/MySQL/UploadDataRepository"
	horariorepositoriorabbit "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio/rabbitMQ/repoRabbit"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//SetupRouter is a func which bind each uri with a handler function
func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	horariorepoRMQ := horariorepositoriorabbit.New(constants.AMQPURL)
	horariosrv := horariosrv.New(horariorepoRMQ)
	uploadrepo := uploaddatarepositorymysql.New()
	uploaddata := uploaddata.New(uploadrepo)
	horarioHandler := handlers.NewHTTPHandler(horariosrv, uploaddata)
	r.GET(constants.PING_URL, handlers.Ping)
	r.GET(constants.GET_AVAILABLE_HOURS_URL, horarioHandler.GetAvailableHours)
	r.POST(constants.UPDATE_SCHEDULER_URL, horarioHandler.PostUpdateScheduler)
	r.GET(constants.LIST_DEGREES_URL, horarioHandler.ListDegrees)
	r.GET(constants.LIST_SCHEDULER_ENTRIES_URL, horarioHandler.GetEntries)
	r.GET(constants.GENERATE_ICAL_URL, horarioHandler.GetICS)
	r.POST(constants.UPLOAD_DATA_DEGREES_URL, horarioHandler.UpdateByCSV)

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
