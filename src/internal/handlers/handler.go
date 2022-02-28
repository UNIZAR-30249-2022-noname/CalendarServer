package handlers

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"

type HTTPHandler struct {
	horarioService    ports.SchedulerService
	uploadDataservice ports.UploadDataservice
}

func NewHTTPHandler(horarioService ports.SchedulerService, uploadData ports.UploadDataservice) *HTTPHandler {
	return &HTTPHandler{
		horarioService: horarioService,
	}
}
