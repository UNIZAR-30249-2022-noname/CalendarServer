package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type HorarioService interface {
	GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error)
	//TODO funciones de este service
}
