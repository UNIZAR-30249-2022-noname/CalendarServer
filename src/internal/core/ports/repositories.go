package ports

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type HorarioRepositorio interface {
	GetAvailableHours(domain.Terna) ([]domain.AvailableHours, error)
	SaveEntry(domain.Entry) error
}
