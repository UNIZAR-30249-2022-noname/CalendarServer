package rabbitamqRepository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/auxFuncs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

const (
	ENV_MODE     = "GATEWAY_MODE"
	testExtesion = "_test"
)

type messageQueue struct {
	Body    interface{} `json:"body"`
	Pattern string      `json:"pattern"`
	Id      string      `json:"id"`
}

/*------------------------------------------------------------------------------------------------------*/
type Repository struct {
	ch      *amqp.Channel
	replies chan amqp.Delivery
}

func New(ch *amqp.Channel, queues []string) (*Repository, error) {

	rp := Repository{ch: ch, replies: make(chan amqp.Delivery, 100)}
	checkMode(queues)
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

func checkMode(queues []string) {
	if os.Getenv(ENV_MODE) == constants.TEST_MODE {
		for i := 0; i < len(queues); i++ {
			queues[i] += testExtesion
		}
	}

}

func (rp *Repository) RCPcallJSON(msg interface{}, pattern string) ([]byte, error) {
	//TODO garantizar exclusion mutua
	corrId := auxFuncs.RandomString(10)
	message := messageQueue{Body: &msg, Pattern: pattern, Id: corrId}
	msgJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

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
