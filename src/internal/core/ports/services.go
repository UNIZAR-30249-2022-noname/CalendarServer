package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

//HorarioServie is the interface which provide access to all the
//scheduler services related
type SchedulerService interface {
	//GetAvaiabledHours is a function which returns a set of [AvailableHours]
	//given a completed [Terna] (not null fields)
	GetAvailableHours(terna domain.DegreeSet) ([]domain.AvailableHours, error)
	ListAllDegrees() ([]domain.DegreeDescription, error)
	GetEntries(terna domain.DegreeSet) ([]domain.Entry, error)
	UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error)
	GetICS(terna domain.DegreeSet) (string, error)
}

type UploadDataService interface {
	UpdateByCSV(csv string) (bool, error)
}

type MonitoringService interface {
	Ping() (bool, error)
}

type UsersService interface {
	GetCredentials(username string) (domain.User, error)
}
