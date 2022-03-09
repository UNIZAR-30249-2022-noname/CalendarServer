package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/docs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/monitoring"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	monitoringrepositoryrabbitamq "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ/monitoringRepository"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var rabbitConn connection.Connection

func config() (handlers.HTTPHandler, error) {
	//Conexión con rabbit
	var err error
	rabbitConn, err = connection.New(constants.AMQPURL)
	if err != nil {
		//TODO
	}
	chMonitoring, err := rabbitConn.NewChannel()
	if err != nil {
		//TODO
	}
	monitoringRepo := monitoringrepositoryrabbitamq.New(chMonitoring)

	return handlers.HTTPHandler{
		Monitoring: monitoring.New(monitoringRepo),
	}, nil

}

//SetupRouter is a func which bind each uri with a handler function
func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handler, err := config()
	if err != nil {
		//TODO
	}
	r.GET(constants.PING_URL, handler.Ping)
	r.GET(constants.GET_AVAILABLE_HOURS_URL, handler.GetAvailableHours)
	r.POST(constants.UPDATE_SCHEDULER_URL, handler.PostUpdateScheduler)
	r.GET(constants.LIST_DEGREES_URL, handler.ListDegrees)
	r.GET(constants.LIST_SCHEDULER_ENTRIES_URL, handler.GetEntries)
	r.GET(constants.GENERATE_ICAL_URL, handler.GetICS)
	r.POST(constants.UPLOAD_DATA_DEGREES_URL, handler.UpdateByCSV)

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
	defer rabbitConn.Disconnect()

	r.Run(":8080")
}
