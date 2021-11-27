package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

//HorarioRepositorio is the interface which provide access to all
//scheduler data related
type HorarioRepositorio interface {
	GetAvailableHours(domain.Terna) ([]domain.AvailableHours, error)
	CreateNewEntry(domain.Entry) error
	DeleteEntry(domain.Entry) error
	ListAllDegrees() ([]domain.DegreeDescription, error)
	DeleteAllEntries() error
}
