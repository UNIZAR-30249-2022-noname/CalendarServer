package rabbitamqRepository

import (
	"encoding/json"
	"os"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/auxFuncs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
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
type errorMessageQueue struct {
	Err        string `json:"err"`
	Message    string `json:"message"`
	IsDisposed bool   `json:"isDisposed"`
}

type DataMessageQueue[T any] struct {
	Response struct {
		Result T `json:"resultado"`
	} `json:"response"`
}

type DataMessageQueueDownload struct {
	Response struct {
		Result struct {
			Type string `json:"type"`
			Data []byte `json:"data"`
		} `json:"resultado"`
	} `json:"response"`
}

/*------------------------------------------------------------------------------------------------------*/
type Repository struct {
	rabbitConn connection.Connection
}

func New(rabbitConn connection.Connection, queues []string) (*Repository, error) {
	rp := Repository{rabbitConn: rabbitConn}
	checkMode(queues)
	ch, _ := rabbitConn.NewChannel()
	defer ch.Close()
	for _, queue := range queues {
		//crear cola de peticiones
		err := connect.PrepareChannel(ch, queue)
		if err != nil {
			return &Repository{}, err
		}
	}
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
	ch, _ := rp.rabbitConn.NewChannel()
	defer ch.Close()
	corrId := auxFuncs.RandomString(10)
	message := messageQueue{Body: &msg, Pattern: pattern, Id: corrId}
	msgJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	//enviar la petición
	err = ch.Publish(
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
	//canal por el que se recibe la respuesta
	msgs, err := ch.Consume(
		constants.REPLY, // queue
		"",              // consumer
		false,           // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)

	if err != nil {
		return nil, err
	}

	for resp := range msgs {
		if corrId == resp.CorrelationId {
			data = resp.Body
			resp.Ack(false)
			break
		}
	}
	errorMsg := errorMessageQueue{}
	json.Unmarshal(data, &errorMsg)
	if errorMsg.Err != "" {
		return nil, apperrors.ErrInternal
	}
	return data, err
}
