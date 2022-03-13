package space

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

type SpaceServiceImp struct {
	spaceRepository ports.SpaceRepository
}

func New(spaceRepository ports.SpaceRepository) *SpaceServiceImp {
	return &SpaceServiceImp{spaceRepository: spaceRepository}
}

func (svc *SpaceServiceImp) Reserve() (bool, error) {
	return svc.spaceRepository.Reserve()
}

func (svc *SpaceServiceImp) ReserveBatch() (bool, error) {
	return svc.spaceRepository.ReserveBatch()
}
