package horariosrv

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

type service struct {
	horarioRepositorio ports.HorarioRepositorio
}

func New(horarioRepositorio ports.HorarioRepositorio) *service {
	return &service{horarioRepositorio: horarioRepositorio}
}

func (srv *service) GetAvaiableHours(terna domain.Terna) ([]domain.AvaiableHours, error) {
	return nil, nil
}
