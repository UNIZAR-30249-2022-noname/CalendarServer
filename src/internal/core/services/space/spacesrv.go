package space

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

type SpaceServiceImp struct {
	spaceRepository ports.SpaceRepository
}

func New(spaceRepository ports.SpaceRepository) *SpaceServiceImp {
	return &SpaceServiceImp{spaceRepository: spaceRepository}
}

func (svc *SpaceServiceImp) Reserve(sp domain.Space, init, end domain.Hour, date string) (string, error) {
	return svc.spaceRepository.Reserve(sp, init, end, date)
}

func (svc *SpaceServiceImp) ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string) (string, error) {
	return svc.spaceRepository.ReserveBatch(spaces, init, end, dates)
}
func (svc *SpaceServiceImp) FilterBy(domain.SpaceFilterParams) (domain.Spaces, error)
