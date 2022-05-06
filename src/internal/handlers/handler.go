package handlers

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"

type HTTPHandler struct {
	Scheduler  ports.SchedulerService
	UploadData ports.UploadDataService
	Monitoring ports.MonitoringService
	Users      ports.UsersService
	Spaces     ports.SpacesService
	Issues     ports.IssueService
}
