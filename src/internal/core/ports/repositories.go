package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type HorarioRepositorio interface {
	//TODO funciones de este repo
	GetAvailableHours(domain.Terna) ([]domain.AvailableHours, error)
}
