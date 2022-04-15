package space

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
)

type SpaceServiceImp struct {
	spaceRepository ports.SpaceRepository
}

func New(spaceRepository ports.SpaceRepository) *SpaceServiceImp {
	return &SpaceServiceImp{spaceRepository: spaceRepository}
}

func (svc *SpaceServiceImp) RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error) {
	if req.Name != "" && req.Date != "" {
		return svc.spaceRepository.RequestInfoSlots(req)
	}
	return domain.AllInfoSlot{}, apperrors.ErrInvalidInput
}

func (svc *SpaceServiceImp) Reserve(sp string, init, end domain.Hour, date, person, event string) (string, error) {
	return svc.spaceRepository.Reserve(sp, init, end, date, person, event)
}

func (svc *SpaceServiceImp) ReserveBatch(spaces []string, init, end domain.Hour, dates []string, person string) (string, error) {
	return svc.spaceRepository.ReserveBatch(spaces, init, end, dates, person)
}
func (svc *SpaceServiceImp) FilterBy(params domain.SpaceFilterParams) ([]domain.Space, error) {
	return svc.spaceRepository.FilterBy(params)
}

func (svc *SpaceServiceImp) CancelReserve(key string) error {

	return svc.spaceRepository.CancelReserve(key)

}

func (svc *SpaceServiceImp) GetReservesOwner(owner string) ([]domain.Reserve, error) {

	return svc.spaceRepository.GetReservesOwner(owner)

}
