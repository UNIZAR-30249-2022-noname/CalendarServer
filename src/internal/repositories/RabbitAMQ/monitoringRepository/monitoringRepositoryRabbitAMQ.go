package monitoringrepositoryrabbitamq

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/auxFuncs"
	"github.com/streadway/amqp"
)

type HorarioRepositorioRabbit struct {
	ch *amqp.Channel
}

func New(ch *amqp.Channel) *HorarioRepositorioRabbit {
	return &HorarioRepositorioRabbit{ch}
}

func (repo *HorarioRepositorioRabbit) Ping() (bool, error) {
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
			CorrelationId: auxFuncs.RandomString(10),
			Body:          []byte("Hola, esto es una prueba, Guapo el que lo lea"),
		})
	return true, err
}

