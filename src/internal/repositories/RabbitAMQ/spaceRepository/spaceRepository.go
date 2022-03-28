package spacerepositoryrabbitamq

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type SpaceRepository struct {
	*rabbitamqRepository.Repository
}

func New(ch *amqp.Channel) (*SpaceRepository, error) {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(ch, queues)
	if err != nil {
		return &SpaceRepository{}, err
	}
	return &SpaceRepository{rp}, nil
}

func (repo *SpaceRepository) RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error) {
	var allInfo domain.AllInfoSlot
	allInfoJSON, err := repo.RCPcallJSON(constants.REQUEST)
	if err != nil {
		return domain.AllInfoSlot{}, err
	}
	json.Unmarshal(allInfoJSON, &allInfo)
	return allInfo, nil

}

func (repo *SpaceRepository) Reserve(space domain.Space, init, end domain.Hour, date, person string) (string, error) {

	return "1", nil
}

func (repo *SpaceRepository) ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string, person string) (string, error) {

	return "0", nil
}

func (repo *SpaceRepository) FilterBy(domain.SpaceFilterParams) ([]domain.Spaces, error) {

	return []domain.Spaces{}, nil
}
