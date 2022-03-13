package spaceRepository

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/auxFuncs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type SpaceRepository struct {
	ch *amqp.Channel
}

func New(ch *amqp.Channel) *SpaceRepository {
	return &SpaceRepository{ch}
}

func (repo *SpaceRepository) Reserve(sp domain.Space, init, end domain.Hour) (bool, error) {
	err := connect.PrepareChannel(repo.ch, constants.RESERVE)
	if err != nil {
		return false, err
	}
	msg, err := json.Marshal(domain.Reserve{Space: sp, Init: init, End: end})
	if err != nil {
		return false, err
	}
	err = repo.ch.Publish(
		"",          // exchange
		constants.RESERVE, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: auxFuncs.RandomString(10),
			Body:          msg,
		})
	return true, err
}

func (repo *SpaceRepository) ReserveBatch(spaces []domain.Space, init, end domain.Hour) (bool, error) {
	err := connect.PrepareChannel(repo.ch, constants.BATCH)
	if err != nil {
		return false, err
	}
	msg, err := json.Marshal(domain.ReserveBatch{Spaces: spaces, Init: init, End: end})
	if err != nil {
		return false, err
	}
	err = repo.ch.Publish(
		"",          // exchange
		constants.BATCH, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: auxFuncs.RandomString(10),
			Body:          msg,
		})
	return true, err
}
