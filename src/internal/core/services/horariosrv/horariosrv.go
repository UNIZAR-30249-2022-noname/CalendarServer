package horariosrv

import (
	"time"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
)

//HorarioServiceImp is the implemetation of [HorarioService] interface.
type HorarioServiceImp struct {
	horarioRepositorio ports.HorarioRepositorio
}

//New is a function which creates a new [HorarioServiceImp]
func New(horarioRepositorio ports.HorarioRepositorio) *HorarioServiceImp {
	return &HorarioServiceImp{horarioRepositorio: horarioRepositorio}
}

//GetAvaiabledHours is a function which returns a set of [AvailableHours]
//given a completed [Terna] (not null fields)
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

	err = srv.horarioRepositorio.CreateNewEntry(entry)
	if err != nil {
		return "", apperrors.ErrInternal
	}
	return time.Now().Format("02/01/2006"), nil
}

func (srv *HorarioServiceImp) ListAllDegrees() ([]domain.DegreeDescription, error) {
	list, err := srv.horarioRepositorio.ListAllDegrees()
	return list, err
}

func (srv *HorarioServiceImp) UpdateScheduler(entries []domain.Entry) (string, error) {
	var lastMod string
	err := srv.horarioRepositorio.DeleteAllEntries(domain.Terna{}) //TODO pasar la correcta
	if err != nil {
		return "", apperrors.ErrSql
	}

	lastMod = time.Now().Format("02/01/2006")

	for i, e := range entries {
		//add
		date, err := srv.CreateNewEntry(e)
		if err != nil {
			return "", apperrors.ErrSql
		}
		if len(entries)-1 == i {
			lastMod = date
		}

	}
	return lastMod, nil
}
