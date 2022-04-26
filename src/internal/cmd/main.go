package main

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/docs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/issue"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/monitoring"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/space"
	uploaddata "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/uploadData"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/users"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	usersrepositorymemory "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/Memory/usersRepository"
	uploadDatarepositoryrabbitamq "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ/UploadDataRepository"
	issuerepositoryrabbitamq "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ/issueRepository"
	monitoringrepositoryrabbitamq "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ/monitoringRepository"
	spacerepositoryrabbitamq "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ/spaceRepository"

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
		panic(err)
	}
	rabbitConn.PurgeAll()
	if err != nil {
		panic(err)
	}
	chMonitoring, err := rabbitConn.NewChannel()
	if err != nil {
		panic(err)
	}
	monitoringRepo := monitoringrepositoryrabbitamq.New(chMonitoring)
	//TODO canal
	chSpace, err := rabbitConn.NewChannel()
	if err != nil {
		//TODO
	}
	//spaceRepoAMQ, _ := spacerepositoryrabbitamq.New(chSpaces)
	spaceRepo, err := spacerepositoryrabbitamq.New(chSpace)
	if err != nil {
		//TODO
	}
	usersRepo := usersrepositorymemory.New()
	chIssues, err := rabbitConn.NewChannel()
	if err != nil {
		//TODO
	}
	issuesRepo, err := issuerepositoryrabbitamq.New(chIssues)
	if err != nil {
		//TODO
	}
	chUploadData, err := rabbitConn.NewChannel()
	uploadDataRepo, err := uploadDatarepositoryrabbitamq.New(chUploadData)
	if err != nil {
		//TODO
	}

	return handlers.HTTPHandler{
		Monitoring: monitoring.New(monitoringRepo),
		Users:      users.New(usersRepo),
		Spaces:     space.New(spaceRepo),
		Issues:     issue.New(issuesRepo),
		UploadData: uploaddata.New(uploadDataRepo),
	}, nil

}

func CorsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "X-Requested-With") 
	return cors.New(config)
}

//SetupRouter is a func which bind each uri with a handler function
func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(CorsConfig())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handler, err := config()
	if err != nil {
		//TODO
	}
	r.GET(constants.PING_URL, handler.Ping)
	r.GET(constants.LOGIN, handler.Login)
	r.GET(constants.FILTER_SPACES, handler.FilterBy)
	r.GET(constants.REQUEST_INFO_SLOTS, handler.RequestInfoSlots)
	r.POST(constants.RESERVE_SPACE, handler.Reserve)
	r.POST(constants.UPLOAD_DATA_DEGREES_URL, handler.UpdateByCSV)
	r.GET(constants.RESERVE_BATCH, handler.ReserveBatch)

	r.GET(constants.CANCEL_RESERVE, handler.CancelReserve)
	r.GET(constants.DELETE_ISSUE, handler.DeleteIssue)
	r.POST(constants.CREATE_ISSUE, handler.CreateIssue)
	r.GET(constants.MODIFY_ISSUE, handler.ChangeStateIssue)
	r.GET(constants.GET_ALL_ISSUES, handler.GetAllIssues)
	r.GET(constants.DOWNLOAD_ISSUES, handler.DownloadIssues)
	r.GET(constants.GET_RESERVES_USER, handler.GetReservesOwner)

	return r
}

func main() {
	// · Swagger · //
	docs.SwaggerInfo.Title = "API UNIZAR calendar and schedule"
	docs.SwaggerInfo.Description = "This is API for managing and visulizating the calendar and schedule of Unizar."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := SetupRouter()
	defer rabbitConn.Disconnect()

	r.Run(":8080")
}
