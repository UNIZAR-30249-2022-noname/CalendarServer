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
	allInfoJSON, err := repo.RCPcallJSON(req, constants.REQINFOSLOT)
	if err != nil {
		return domain.AllInfoSlot{}, err
	}
	json.Unmarshal(allInfoJSON, &allInfo)
	return allInfo, nil

}

func (repo *SpaceRepository) Reserve(space string, init, end domain.Hour, date, person string) (string, error) {
	var reserveId string
	//TODO Añadir campo evento y KEY
	reserveIdJSON, err := repo.RCPcallJSON(domain.Reserve{Space: space, Day: date, 
		Event: "Evento", Scheduled: []domain.Hour{init,end}, Owner: person, Key: "0",}, constants.RESERVE)
	if err != nil {
		return "", err
	}
	json.Unmarshal(reserveIdJSON, &reserveId)

	return reserveId, nil
}

func (repo *SpaceRepository) ReserveBatch(spaces []string, init, end domain.Hour, dates []string, person string) (string, error) {

	return "0", nil
}

func (repo *SpaceRepository) FilterBy(spaceParams domain.SpaceFilterParams) ([]domain.Space, error) {
	var spaces []domain.Space
	spacesJSON, err := repo.RCPcallJSON(spaceParams, constants.SPFILTER)
	if err != nil {
		return []domain.Space{}, err
	}
	json.Unmarshal(spacesJSON, &spaces)

	return spaces, nil
}

//TODO 
/*
func (repo *SpaceRepository) UploadSpaces() (string, error) {
	message := "hola"
	messageUploadSpace := &domain.MessageQueueUploadSpace{
		Pattern:   "importar-espacios",
		Id:        auxFuncs.RandomString(5),
		Fichero: message,
	}
	_, err := repo.RCPcallJSON(messageUploadSpace, messageUploadSpace.Id)
	if err != nil {
		return "", err
	}

	return message, nil
}
*/