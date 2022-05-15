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
	UpdateByCSV(csv, privileges string) (bool, error)
}

type MonitoringService interface {
	Ping() (bool, error)
}

type UsersService interface {
	GetCredentials(username string) (domain.User, error)
}

type SpacesService interface {
	FilterBy(domain.SpaceFilterParams) ([]domain.Space, error)
	RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error)
	Reserve(sp string, init, end domain.Hour, date, person, event string) (string, error)
	ReserveBatch(spaces []string, init, end domain.Hour, dates []string, person string) (string, error)
	CancelReserve(key string) error
	GetReservesOwner(owner string) ([]domain.Reserve, error)
}

type IssueService interface {
	GetAll() ([]domain.Issue, error)
	Delete(key string) error
	Create(issue domain.Issue) error
	ChangeState(key string, state int) error
	DownloadIssues(building string) ([]byte ,error)
}
