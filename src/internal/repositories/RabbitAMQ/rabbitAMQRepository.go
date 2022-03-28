package rabbitamqRepository

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/auxFuncs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type Repository struct {
	ch      *amqp.Channel
	replies chan amqp.Delivery
}

func New(ch *amqp.Channel, queues []string) (*Repository, error) {
	rp := Repository{ch: ch, replies: make(chan amqp.Delivery, 100)}
	for _, queue := range queues {
		//crear cola de peticiones
		err := connect.PrepareChannel(rp.ch, queue)
		if err != nil {
			return &Repository{}, err
		}
	}
	//canal por el que se recibe la respuesta
	msgs, err := rp.ch.Consume(
		constants.REPLY, // queue
		"",              // consumer
		false,           // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		return &Repository{}, err
	}
	go func() {
		for resp := range msgs {
			rp.replies <- resp
		}

	}()
	return &rp, nil
}

func (rp *Repository) RCPcallJSON(msg interface{}) ([]byte, error) {
	//TODO garantizar exclusiçon mutua
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	corrId := auxFuncs.RandomString(10)

	//enviar la petición
	err = rp.ch.Publish(
		"",                // exchange
		constants.REQUEST, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			Body:          msgJSON,
			ReplyTo:       constants.REPLY,
		})
	var data []byte
	for resp := range rp.replies {
		if corrId == resp.CorrelationId {
			data = resp.Body
			break
		} else {
			rp.replies <- resp
		}
	}
	return data, err
}
