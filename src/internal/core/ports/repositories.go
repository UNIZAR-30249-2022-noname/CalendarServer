package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

//HorarioRepositorio is the interface which provide access to all
//scheduler data related
type HorarioRepositorio interface {
	GetAvailableHours(domain.Terna) ([]domain.AvailableHours, error)
	CreateNewEntry(domain.Entry) error
	DeleteEntry(domain.Entry) error
	ListAllDegrees() ([]domain.DegreeDescription, error)
	DeleteAllEntries(terna domain.Terna) error
	GetEntries(domain.Terna) ([]domain.Entry, error)
	CreateNewDegree(id int, name string) (bool, error)
	CreateNewSubject(id int, name string, degreeCode int) (bool, error)
	CreateNewYear(year int, degreeCode int) (bool, error)
	CreateNewGroup(group int, yearCode int) (bool, error)
	CreateNewHour(available, total, subjectCode, groupCode, kind int, group, week string) (bool, error)
	RawExec(exec string) error
}
