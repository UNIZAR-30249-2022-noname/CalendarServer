package horariosrv

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

//service is the implemetation of [HorarioService] interface.
type service struct {
	horarioRepositorio ports.HorarioRepositorio
}

//New is a function which creates a new [service]
func New(horarioRepositorio ports.HorarioRepositorio) *service {
	return &service{horarioRepositorio: horarioRepositorio}
}

//GetAvaiabledHours is a function which returns a set of [AvailableHours]
//given a completed [Terna] (not null fields)
func (srv *service) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {
	res, err := srv.horarioRepositorio.GetAvailableHours(terna)
	if err != nil {
		return []domain.AvailableHours{}, err
	}

	return res, nil
}

func (srv *service) CreateNewEntry(entry domain.Entry) (string, error) {
	return "", nil
}
