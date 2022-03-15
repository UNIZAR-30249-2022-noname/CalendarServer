package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

//SchedulerRepository is the interface which provide access to all
//scheduler data related
type SchedulerRepository interface {
	GetAvailableHours(domain.DegreeSet) ([]domain.AvailableHours, error)
	CreateNewEntry(domain.Entry) error
	DeleteEntry(domain.Entry) error
	ListAllDegrees() ([]domain.DegreeDescription, error)
	DeleteAllEntries(terna domain.DegreeSet) error
	GetEntries(domain.DegreeSet) ([]domain.Entry, error)
}

type UploadDataRepository interface {
	CreateNewDegree(id int, name string) (bool, error)
	CreateNewSubject(id int, name string, degreeCode int) (bool, error)
	CreateNewYear(year int, degreeCode int) (bool, error)
	CreateNewGroup(group int, yearCode int) (bool, error)
	CreateNewHour(available, total, subjectCode, groupCode, kind int, group, week string) (bool, error)
	RawExec(exec string) error
}

type MonitoringRepository interface {
	Ping() (bool, error)
}

type UsersRepository interface {
	GetCredentials(username string) (domain.User, error)
}

type SpaceRepository interface {
	FilterBy(domain.SpaceFilterParams) ([]domain.Spaces, error)
	Reserve(sp domain.Space, init, end domain.Hour, date string) (string, error)
	ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string) (string, error)
}
