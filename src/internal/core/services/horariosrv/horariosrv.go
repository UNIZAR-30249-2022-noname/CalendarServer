package horariosrv

import (
	"time"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
)

type HorarioServiceImp struct {
	horarioRepositorio ports.HorarioRepositorio
}

func New(horarioRepositorio ports.HorarioRepositorio) *HorarioServiceImp {
	return &HorarioServiceImp{horarioRepositorio: horarioRepositorio}
}

func (srv *HorarioServiceImp) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {
	res, err := srv.horarioRepositorio.GetAvailableHours(terna)
	if err != nil {
		return []domain.AvailableHours{}, err
	}

	return res, nil
}

func (srv *HorarioServiceImp) CreateNewEntry(entry domain.Entry) (string, error) {
	err := entry.IsValid()
	if err != nil {
		return "", err
	}

	//Check if the entry has valid time interval
	if entry.Init.IsLaterThan(entry.End) {
		return "", apperrors.ErrInvalidInput
	}

	err = srv.horarioRepositorio.SaveEntry(entry)
	if err != nil {
		return "", apperrors.ErrInternal
	}
	return time.Now().Format("02/01/2006"), nil
}
