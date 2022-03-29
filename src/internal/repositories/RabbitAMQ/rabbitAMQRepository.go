package rabbitamqRepository

import (
	"encoding/json"
	"fmt"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/auxFuncs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

/**BASURA QUE HE CREADO PARA PROBAR HASTA LA LINEA 34*/
type MessageQueue struct {
	Body    string `json:"body"`
	Pattern string `json:"pattern"`
	Age     string `json:"age"`
	Id      string `json:"id"`
}

func NewMessageQueue(body string, pattern string, age string, id string) *MessageQueue {
	return &MessageQueue{
		body, pattern, age, id,
	}
}

func (m *MessageQueue) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(m)

	if err != nil {
		return nil, err
	}
	return bytes, err
}

/*------------------------------------------------------------------------------------------------------*/
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
		true,            // auto-ack
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

func (rp *Repository) RCPcallJSON(msg interface{}, msgId string) ([]byte, error) {
	//TODO garantizar exclusion mutua
	message := NewMessageQueue("10:00", "realizar-reserva", "Sergio", "1234")
	// TODO: check the error
	bodyFake, err := message.Marshal()
	//msgJSON, err := json.Marshal(msg)
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
			Body:          bodyFake,
			MessageId: 	   msgId,
			ReplyTo:       constants.REPLY,
		})
	var data []byte
	for resp := range rp.replies {
		fmt.Println(resp.Body)
		myString := string(resp.Body[:])
		fmt.Println(myString)
		if corrId == resp.CorrelationId {
			data = resp.Body
			break
		} else {
			rp.replies <- resp
		}
	}
	return data, err
}
