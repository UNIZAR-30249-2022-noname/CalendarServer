package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type HorarioService interface {
	GetAvaiableHours(terna domain.Terna) ([]domain.AvaiableHours, error)
	//TODO funciones de este service
}
