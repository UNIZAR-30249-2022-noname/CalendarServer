package spacerepositorymemory

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
)

type SpaceRepository struct {
}

func New() *SpaceRepository {
	return &SpaceRepository{}
}

//TODO Este sobra
func (repo *SpaceRepository) Reserve(sp domain.Space, init, end domain.Hour, date string) (string, error) {

	return "1", nil
}

func (repo *SpaceRepository) ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string) (string, error) {

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
