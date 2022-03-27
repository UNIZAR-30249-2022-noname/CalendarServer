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
		sd := domain.SlotData{
			Name: "A1",
			Capacity: 5,
			Description: "Lorem ipsum no leas mas porque esto es dummy text",
			Building: "Ada",
			Floor: "baja",
			Type: "aula",
		  };

		  
		  is := []domain.InfoSlots{
			{
			Hour: 8,
			Occupied: false,
			},
			{
				Hour: 9,
				Occupied: true,
				Person: "Urrikote",
			},
			{
				Hour: 10,
				Occupied: false,
			},
			{
				Hour: 11,
				Occupied: false,
			},
			{
				Hour: 12,
				Occupied: true,
				Person: "Urrikyu",
			},
			{
				Hour: 13,
				Occupied: false,
			},
			{
				Hour: 14,
				Occupied: false,
			},
			{
				Hour: 15,
				Occupied: true,
				Person: "Urriuuuu",
			},
			{
				Hour: 16,
				Occupied: false,
			},
			{
				Hour: 17,
				Occupied: false,
			},
			{
				Hour: 8,
				Occupied: true,
				Person: "Urrikoncio",
			},
			{
				Hour: 19,
				Occupied: false,
			},
			{
				Hour: 20,
				Occupied: false,
			},
		}
	
		allInfo := domain.AllInfoSlot{
			SlotData: sd,
			InfoSlots: is,
		}

	return allInfo, nil
}

func (repo *SpaceRepository) Reserve(sp domain.Space, init, end domain.Hour, date, person string) (string, error) {

	return "1", nil
}

func (repo *SpaceRepository) ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string, person string) (string, error) {

	return "1", nil
}

func (repo *SpaceRepository) FilterBy(domain.SpaceFilterParams) ([]domain.Spaces, error) {

	return []domain.Spaces{
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
