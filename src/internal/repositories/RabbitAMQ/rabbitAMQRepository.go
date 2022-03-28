package rabbitamqRepository

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/auxFuncs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type Repository struct {
	ch *amqp.Channel
}

func New(ch *amqp.Channel, queues []string) (*Repository, error) {
	rp := Repository{ch: ch}
	for _, queue := range queues {
		err := connect.PrepareChannel(rp.ch, queue)
		err = connect.PrepareChannel(rp.ch, queue+constants.REPPLY_EXTENSION)
		if err != nil {
			return &Repository{}, err
		}

	}
	return &rp, nil

}

func (rp *Repository) RCPcallJSON(queue string, msg interface{}) (interface{}, error) {

	//canal por el que se recibe la respuesta
	msgs, err := rp.ch.Consume(
		constants.BATCH_REPLY, // queue
		"",                    // consumer
		false,                 // auto-ack
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                   // args
	)

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	corrId := auxFuncs.RandomString(10)

	var result interface{}
	//enviar la petición
	err = rp.ch.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			Body:          msgJSON,
			ReplyTo:       queue + constants.REPPLY_EXTENSION,
		})

	for resp := range msgs {
		if corrId == resp.CorrelationId {
			if err == nil {
				json.Unmarshal(resp.Body, result)
			}
			break
		}
	}
	return result, err
}
