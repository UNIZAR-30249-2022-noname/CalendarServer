package spacerepositorymemory

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
)

type SpaceRepository struct {
}

func New() *SpaceRepository {
	return &SpaceRepository{}
}

func (repo *SpaceRepository) RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error) {
	sd := domain.Space{
		Name:        "A1",
		Capacity:    5,
		Description: "Lorem ipsum no leas mas porque esto es dummy text",
		Building:    "Ada",
		Floor:       "baja",
		Kind:        "aula",
	}

	is := []domain.InfoSlots{
		{
			Hour:     8,
			Occupied: false,
		},
		{
			Hour:     9,
			Occupied: true,
			Person:   "Urrikote",
		},
		{
			Hour:     10,
			Occupied: false,
		},
		{
			Hour:     11,
			Occupied: false,
		},
		{
			Hour:     12,
			Occupied: true,
			Person:   "Urrikyu",
		},
		{
			Hour:     13,
			Occupied: false,
		},
		{
			Hour:     14,
			Occupied: false,
		},
		{
			Hour:     15,
			Occupied: true,
			Person:   "Urriuuuu",
		},
		{
			Hour:     16,
			Occupied: false,
		},
		{
			Hour:     17,
			Occupied: false,
		},
		{
			Hour:     8,
			Occupied: true,
			Person:   "Urrikoncio",
		},
		{
			Hour:     19,
			Occupied: false,
		},
		{
			Hour:     20,
			Occupied: false,
		},
	}

	allInfo := domain.AllInfoSlot{
		SlotData:  sd,
		InfoSlots: is,
	}

	return allInfo, nil
}

func (repo *SpaceRepository) Reserve(sp string, init, end domain.Hour, date, person string) (string, error) {

	return "1", nil
}

func (repo *SpaceRepository) ReserveBatch(spaces []string, init, end domain.Hour, dates []string, person string) (string, error) {

	return "1", nil
}

func (repo *SpaceRepository) FilterBy(domain.SpaceFilterParams) ([]domain.Space, error) {

	return []domain.Space{
		{
			Name:     "A1",
			Capacity: 20,
			Building: "Ada",
			Kind:     "aula",
		},
		{
			Name:     "A2",
			Capacity: 30,
			Building: "Ada",
			Kind:     "aula",
		},
		{
			Name:     "L0",
			Capacity: 35,
			Building: "Ada",
			Kind:     "laboratorio",
		},
	}, nil

}

func (repo *SpaceRepository) CancelReserve(key string) error {
	return nil
}

func (repo *SpaceRepository) GetReservesOwner(owner string) ([]domain.Reserve, error) {
	return []domain.Reserve{
		{
			Space: "A0.11",
			Day:   "12/2/2022",
			Event: "Prog 1",
			Scheduled: []domain.Hour{
				{Hour: 9, Min: 00},
				{Hour: 10, Min: 00},
			},
			Owner: "Luigui",
			Key:   "1",
		},
		{
			Space: "A0.13",
			Day:   "12/2/2022",
			Event: "Prog 2",
			Scheduled: []domain.Hour{
				{Hour: 10, Min: 00},
				{Hour: 11, Min: 00},
			},
			Owner: "Luigui",
			Key:   "2",
		},
		{
			Space: "A1.13",
			Day:   "12/2/2022",
			Event: "IC",
			Scheduled: []domain.Hour{
				{Hour: 11, Min: 00},
				{Hour: 12, Min: 00},
			},
			Owner: "Luigui",
			Key:   "3",
		},
	}, nil

}
