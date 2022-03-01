package horariosrv

import (
	"time"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
)

//SchedulerServiceImp is the implemetation of [SchedulerService] interface.
type SchedulerServiceImp struct {
	horarioRepositorioRMQ ports.RabbitRepository
}

//New is a function which creates a new [SchedulerServiceImp]
func New(horarioRepositorioRMQ ports.RabbitRepository) *SchedulerServiceImp {
	return &SchedulerServiceImp{horarioRepositorioRMQ: horarioRepositorioRMQ}
}


func (srv *SchedulerServiceImp) Monitoring() (bool, error) {
	return true, nil
}

//GetAvaiabledHours is a function which returns a set of [AvailableHours]
//given a completed [Terna] (not null fields)
func (srv *SchedulerServiceImp) GetAvailableHours(terna domain.DegreeSet) ([]domain.AvailableHours, error) {
	return []domain.AvailableHours{}, apperrors.ErrToDo
}

func (srv *SchedulerServiceImp) CreateNewEntry(entry domain.Entry) (string, error) {
	err := entry.IsValid()
	if err != nil {
		return "", err
	}

	//Check if the entry has valid time interval
	if entry.Init.IsLaterThan(entry.End) {
		return "", apperrors.ErrInvalidInput
	}

	return time.Now().Format("02/01/2006"), apperrors.ErrToDo
}

func (srv *SchedulerServiceImp) ListAllDegrees() ([]domain.DegreeDescription, error) {

	return []domain.DegreeDescription{}, apperrors.ErrToDo
}

func (srv *SchedulerServiceImp) UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error) {
	lastMod := time.Now().Format("02/01/2006")
	return lastMod, apperrors.ErrToDo
}

func (srv *SchedulerServiceImp) GetEntries(terna domain.DegreeSet) ([]domain.Entry, error) {

	if terna.Degree == "" || terna.Year == 0 || terna.Group == "" {
		return []domain.Entry{}, apperrors.ErrInvalidInput
	}
	
	return []domain.Entry{}, apperrors.ErrToDo

}

func (srv *SchedulerServiceImp) GetICS(terna domain.DegreeSet) (string, error) {
	if terna.Degree == "" || terna.Year == 0 || terna.Group == "" {
		return "", apperrors.ErrInvalidInput
	}
	/*
	entries, err := srv.horarioRepositorio.GetEntries(terna)
	if err != nil {
		return "", apperrors.ErrSql
	}
	cal := ics.NewCalendar()
	t := time.Now()
	month := t.Month()
	year := t.Year()
	if month < 8 {
		month = 2
	} else {
		month = 9
	}
	taux := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	for i, entry := range entries {
		event := cal.AddEvent(fmt.Sprintf("%d@unizar.es", i))
		event.SetSummary(entry.Subject.Name)
		day := (8-int(taux.Weekday()))%7 + entry.Weekday + 1
		t1 := time.Date(year, month, day, entry.Init.Hour, entry.Init.Min, 0, 0, t.Location())
		event.SetStartAt(t1)
		t2 := time.Date(year, month, day, entry.End.Hour, entry.End.Min, 0, 0, t.Location())
		event.SetEndAt(t2)
		event.AddRrule("FREQ=DAILY;INTERVAL=7")
		i++
	}
	
	return cal.Serialize(), nil
	*/
	return "", apperrors.ErrToDo
}
