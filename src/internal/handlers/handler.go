package handlers

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"

type HTTPHandler struct {
	SchedulerService  ports.SchedulerService
	UploadDataService ports.UploadDataservice
	Monitoring        ports.Monitoring
}
