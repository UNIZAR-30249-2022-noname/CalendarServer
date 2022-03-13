package spaceRepository

import (
	"math/rand"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/streadway/amqp"
)

type SpaceRepository struct {
	ch *amqp.Channel
}

func New(ch *amqp.Channel) *SpaceRepository {
	return &SpaceRepository{ch}
}

func (repo *SpaceRepository) Reserve() (bool, error) {
	q, err := repo.ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	_ = q
	if err != nil {
		return false, apperrors.ErrConn
	}
	err = repo.ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return false, apperrors.ErrConn
	}
	err = repo.ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: RandomString(10),
			Body:          []byte("Hola, esto es una prueba, Guapo el que lo lea"),
		})
	return true, err
}

func (repo *SpaceRepository) ReserveBatch() (bool, error) {
	q, err := repo.ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	_ = q
	if err != nil {
		return false, apperrors.ErrConn
	}
	err = repo.ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return false, apperrors.ErrConn
	}
	err = repo.ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: RandomString(10),
			Body:          []byte("Hola, esto es una prueba, Guapo el que lo lea"),
		})
	return true, err
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
