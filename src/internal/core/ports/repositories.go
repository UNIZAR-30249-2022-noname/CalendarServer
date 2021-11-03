package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

//HorarioRepositorio is the interface which provide access to all
//scheduler data related
type HorarioRepositorio interface {
	//TODO funciones de este repo
	GetAvailableHours(domain.Terna) ([]domain.AvailableHours, error)
}
