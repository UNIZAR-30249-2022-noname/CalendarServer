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

//TODO Poner fecha
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


//TODO Poner fechas
func (repo *SpaceRepository) ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string) (bool, error) {

	msg, err := json.Marshal(domain.ReserveBatch{Spaces: spaces, Init: init, End: end, Dates: dates})
	if err != nil {
		return false, err
	}

	msgs, err := repo.ch.Consume(
		constants.BATCH_REPLY, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	corrId := auxFuncs.RandomString(10)
	err = repo.ch.Publish(
		"",          // exchange
		constants.BATCH, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo: constants.BATCH_REPLY,
			Body:          msg,
		})
	if err != nil {
		return false, err
	}
	
	lastId := "-1"
	for resp := range msgs {
		if corrId == resp.CorrelationId {
			if err == nil {
				json.Unmarshal(resp.Body, &lastId)
			}
			break
		}
	}
	return true, err
}
