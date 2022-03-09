package monitoring

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

//MonitoringServiceImp is the implemetation of [SchedulerService] interface.
type MonitoringServiceImp struct {
	monitoringRepository ports.MonitoringRepository
}

//New is a function which creates a new [SchedulerServiceImp]
func New(monitoringRepository ports.MonitoringRepository) *MonitoringServiceImp {
	return &MonitoringServiceImp{monitoringRepository: monitoringRepository}
}

func (srv *MonitoringServiceImp) Monitoring() (bool, error){
	return srv.monitoringRepository.Monitoring()
}