package spaceRepository

import (
	"time"

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

func (repo *SpaceRepository) Reserve(sp domain.Space, init, end time.Time) (bool, error) {
	err := connect.PrepareChannel(repo.ch, constants.RESERVE)
	if err != nil {
		return false, err
	}
	err = repo.ch.Publish(
		"",          // exchange
		"reserve_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: auxFuncs.RandomString(10),
			Body:          []byte("Hola, esto es una prueba, Guapo el que lo lea"),
		})
	return true, err
}

func (repo *SpaceRepository) ReserveBatch() (bool, error) {
	err := connect.PrepareChannel(repo.ch, constants.BATCH)
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
			Body:          []byte("Hola, esto es una prueba, Guapo el que lo lea"),
		})
	return true, err
}
