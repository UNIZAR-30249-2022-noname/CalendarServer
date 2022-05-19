package scheduler

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
)

//SchedulerServiceImp is the implemetation of [SchedulerService] interface.
type SchedulerServiceImp struct {
	schedluerRepository ports.SchedulerRepository
}

//New is a function which creates a new [SchedulerServiceImp]
func New(schedluerRepository ports.SchedulerRepository) *SchedulerServiceImp {
	return &SchedulerServiceImp{schedluerRepository: schedluerRepository}
}

//GetAvaiabledHours is a function which returns a set of [AvailableHours]
//given a completed [Terna] (not null fields)
func (srv *SchedulerServiceImp) GetAvailableHours(terna domain.DegreeSet) ([]domain.AvailableHours, error) {
	return srv.GetAvailableHours(terna)
}

func (srv *SchedulerServiceImp) CreateNewEntry(entry domain.Entry) (string, error) {
	return "", apperrors.ErrToDo
}

func (srv *SchedulerServiceImp) ListAllDegrees() ([]domain.DegreeDescription, error) {

	return srv.schedluerRepository.ListAllDegrees()
}

func (srv *SchedulerServiceImp) UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error) {
	return srv.schedluerRepository.UpdateScheduler(entries, terna)
}

func (srv *SchedulerServiceImp) GetEntries(terna domain.DegreeSet) ([]domain.Entry, error) {
	return srv.schedluerRepository.GetEntries(terna)

}

func (srv *SchedulerServiceImp) GetICS(terna domain.DegreeSet) (string, error) {
	return "", apperrors.ErrToDo
}
