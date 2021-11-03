package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

//HorarioServie is the interface which provide access to all the
//scheduler services related
type HorarioService interface {
	//GetAvaiabledHours is a function which returns a set of [AvailableHours]
	//given a completed [Terna] (not null fields)
	GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error)
	CreateNewEntry(entry domain.Entry) (string, error)
}
