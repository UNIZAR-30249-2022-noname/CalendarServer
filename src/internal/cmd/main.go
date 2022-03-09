package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/docs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	rabbit "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/monitoring"
	uploaddata "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/uploadData"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	uploaddatarepositorymysql "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio/MySQL/UploadDataRepository"
	horarioRepositorio "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio/MySQL/horarioRepositorio"
	horariorepositoriorabbit "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio/rabbitMQ/repoRabbit"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type services struct {
	scheduler  ports.SchedulerService
	uploadData ports.UploadDataservice
	rabbit 	   ports.RabbitService
}

var rabbitConn connection.Connection

func config() (services, error) {
	//Conexión con rabbit
	var err error
	rabbitConn, err = connection.New(constants.AMQPURL)
	if err != nil {
		//TODO
	}
	chScheduler, err := rabbitConn.NewChannel()
	if err != nil {
		//TODO
	}
	schedulerRepo := horarioRepositorio.New()
	uploadrepo := uploaddatarepositorymysql.New()
	rabbitRepo := horariorepositoriorabbit.New(chScheduler)

	return services{
		scheduler:  horariosrv.New(schedulerRepo),
		uploadData: uploaddata.New(uploadrepo),
		rabbit:  	rabbit.New(rabbitRepo),
	}, nil

}

//SetupRouter is a func which bind each uri with a handler function
func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srvs, err := config()
	if err != nil {
		//TODO
	}
	horarioHandler := handlers.NewHTTPHandler(srvs.scheduler, srvs.uploadData, srvs.rabbit)
	r.GET(constants.PING_URL, horarioHandler.Ping)
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
	defer rabbitConn.Disconnect()

	r.Run(":8080")
}
