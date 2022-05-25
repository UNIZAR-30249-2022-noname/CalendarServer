package spacerepositoryrabbitamq

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
)

type SpaceRepository struct {
	*rabbitamqRepository.Repository
}

func New(rabbitConn connection.Connection) (*SpaceRepository, error) {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(rabbitConn, queues)
	if err != nil {
		return &SpaceRepository{}, err
	}
	return &SpaceRepository{rp}, nil
}

func (repo *SpaceRepository) RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error) {
	var reply rabbitamqRepository.DataMessageQueue[domain.AllInfoSlot]
	replyJSON, err := repo.RCPcallJSON(req, constants.REQINFOSLOT)
	if err != nil {
		return domain.AllInfoSlot{}, err
	}
	json.Unmarshal(replyJSON, &reply)

	return reply.Response.Result, nil

}

func (repo *SpaceRepository) Reserve(space string, init, end domain.Hour, date, person, event string) (string, error) {
	var reserveId string
	//TODO Añadir campo evento y KEY
	reserveIdJSON, err := repo.RCPcallJSON(domain.Reserve{Space: space, Day: date,
		Event: event, Scheduled: []domain.Hour{init, end}, Owner: person, Key: "0"}, constants.RESERVE)
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
	var reply rabbitamqRepository.DataMessageQueue[[]domain.Space]
	replyJSON, err := repo.RCPcallJSON(spaceParams, constants.SPFILTER)
	if err != nil {
		return []domain.Space{}, err
	}
	json.Unmarshal(replyJSON, &reply)

	return reply.Response.Result, nil
}

func (repo *SpaceRepository) CancelReserve(key string) error {
	return apperrors.ErrToDo
}

func (repo *SpaceRepository) GetReservesOwner(owner string) ([]domain.Reserve, error) {

	var reply rabbitamqRepository.DataMessageQueue[[]domain.Reserve]
	replyJSON, err := repo.RCPcallJSON(owner, constants.GET_RESERVES_USER)
	if err != nil {
		return []domain.Reserve{}, err
	}
	json.Unmarshal(replyJSON, &reply)

	return reply.Response.Result, nil

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
