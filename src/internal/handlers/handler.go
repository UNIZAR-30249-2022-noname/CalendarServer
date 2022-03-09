package handlers

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"

type HTTPHandler struct {
	horarioService    ports.SchedulerService
	uploadDataservice ports.UploadDataservice
	rabbit 	   ports.RabbitService
}

func NewHTTPHandler(horarioService ports.SchedulerService, uploadData ports.UploadDataservice, rabbit ports.RabbitService) *HTTPHandler {
	return &HTTPHandler{
		horarioService: horarioService,
		rabbit: rabbit,
	}
}
