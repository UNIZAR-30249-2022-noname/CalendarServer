package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

//SchedulerRepository is the interface which provide access to all
//scheduler data related
type SchedulerRepository interface {
	GetAvailableHours(req domain.DegreeSet) ([]domain.AvailableHours, error)
	UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error)
	DeleteEntry(req domain.Entry) error
	ListAllDegrees() ([]domain.DegreeDescription, error)
	DeleteAllEntries(terna domain.DegreeSet) error
	GetEntries(req domain.DegreeSet) ([]domain.Entry, error)
}

type UploadDataRepository interface {
	UpdateSpacesByCSV(req string) (bool, error)
	UpdateDegreesByCSV(req string) (bool, error)
}

type MonitoringRepository interface {
	Ping() (bool, error)
}

type UsersRepository interface {
	GetCredentials(username string) (domain.User, error)
}

type SpaceRepository interface {
	FilterBy(domain.SpaceFilterParams) ([]domain.Space, error)
	RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error)
	Reserve(sp string, init, end domain.Hour, date, person, event string) (string, error)
	ReserveBatch(spaces []string, init, end domain.Hour, dates []string, person string) (string, error)
	CancelReserve(key string) error
	GetReservesOwner(owner string) ([]domain.Reserve, error)
}

type IssueRepository interface {
	GetAll() ([]domain.Issue, error)
	Delete(key string) error
	Create(issue domain.Issue) error
	ChangeState(key string, state int) error
	DownloadIssues(building string) ([]byte ,error)
}
