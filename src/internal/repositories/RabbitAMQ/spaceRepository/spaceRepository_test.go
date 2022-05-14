package spacerepositoryrabbitamq_test

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	spaceRepo "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ/spaceRepository"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

const (
	ENV_MODE     = "GATEWAY_MODE"
	testExtesion = "_test"
)

func checkMode(queues []string) {
	if os.Getenv(ENV_MODE) == constants.TEST_MODE {
		for i := 0; i < len(queues); i++ {
			queues[i] += testExtesion
		}
	}
}

func TestDeleteQueueBeforeTest(t *testing.T) {
	assert := assert.New(t)
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReqInfo, err := rabbitConn.NewChannel()
	chReqInfo.QueueDelete(constants.REQUEST, true, false, true)
	chReqInfo.QueueDelete(constants.REPLY, true, false, true)
}

func TestRequestInfoSlots(t *testing.T) {
	t.Skip() //remove for activating it
	queues := []string{constants.REQUEST, constants.REPLY}
	checkMode(queues)
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	data := domain.ReqInfoSlot{Name: "A1", Date: s}
	rabbitConn, err := connection.New(constants.AMQPURL)
	rabbitConn.PurgeAll()
	assert.Equal(err, nil, "Shouldn't be an error")
	chReqInfo, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo, err := spaceRepo.New(rabbitConn)

	//Simulated server
	msgs, _ := chReqInfo.Consume(
		queues[0], // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	myResponse := domain.AllInfoSlot{
		SlotData: domain.Space{
			Name:        "A1",
			Capacity:    5,
			Description: "Lorem ipsum no leas mas porque esto es dummy text requestinfoslots",
			Building:    "Ada",
			Floor:       "baja",
			Kind:        "aula",
		},
		InfoSlots: []domain.InfoSlots{
			{
				Hour:     8,
				Occupied: false,
			},
			{
				Hour:     9,
				Occupied: true,
				Person:   "Urrikote",
			},
		},
	}
	corrId := "-1"
	go func() {
		for resp := range msgs {
			corrId = resp.CorrelationId
			response, _ := json.Marshal(myResponse)
			chReqInfo.Publish(
				"",              // exchange
				queues[1], // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
				//resp.Ack(false)
				break
		}
	}()

	//RUN
	done, err := spaceRepo.RequestInfoSlots(data)
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(done, myResponse, "Should be positive")
}

func TestRequestInfoSlotsMultiple(t *testing.T) {
	t.Skip() //remove for activating it
	queues := []string{constants.REQUEST, constants.REPLY}
	checkMode(queues)
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	data := domain.ReqInfoSlot{Name: "A1", Date: s}
	rabbitConn, err := connection.New(constants.AMQPURL)
	rabbitConn.PurgeAll()
	assert.Equal(err, nil, "Shouldn't be an error")
	chReqInfo, err := rabbitConn.NewChannel()
	spaceRepo, err := spaceRepo.New(rabbitConn)

	//Simulated server
	msgs, _ := chReqInfo.Consume(
		queues[0], // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	myResponse := domain.AllInfoSlot{
		SlotData: domain.Space{
			Name:        "A1",
			Capacity:    5,
			Description: "Lorem ipsum no leas mas porque esto es dummy text multiple",
			Building:    "Ada",
			Floor:       "baja",
			Kind:        "aula",
		},
		InfoSlots: []domain.InfoSlots{
			{
				Hour:     8,
				Occupied: false,
			},
			{
				Hour:     9,
				Occupied: true,
				Person:   "Urrikote",
			},
		},
	}
	corrId := "-1"
	go func() {
		i := 0
		for resp := range msgs {
			corrId = resp.CorrelationId
			response, _ := json.Marshal(myResponse)
			chReqInfo.Publish(
				"",              // exchange
				queues[1], // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
			//resp.Ack(false)
			i++
			if (i>3){
			break
			}
		}
	}()

	//RUN
	done, err := spaceRepo.RequestInfoSlots(data)
	done, err = spaceRepo.RequestInfoSlots(data)
	done, err = spaceRepo.RequestInfoSlots(data)
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(done, myResponse, "Should be positive")
}
func TestReserve(t *testing.T) {
	t.Skip() //remove for activating it
	queues := []string{constants.REQUEST, constants.REPLY}
	checkMode(queues)
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	rabbitConn, err := connection.New(constants.AMQPURL)
	rabbitConn.PurgeAll()
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo, _ := spaceRepo.New(rabbitConn)
	msgs, _ := chReserve.Consume(
		queues[0], // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	corrId := "-1"
	go func() {
		for resp := range msgs {
			corrId = resp.CorrelationId
			response, _ := json.Marshal("1")
			chReserve.Publish(
				"",              // exchange
				queues[1], // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
			//resp.Ack(false)
			break
		}
	}()

	done, err := spaceRepo.Reserve("", domain.Hour{Hour: 12, Min: 30}, domain.Hour{Hour: 12, Min: 30}, s, "Iñigol", "")
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(done, "1", "Should be true")
}

func TestReserveBatch(t *testing.T) {
	t.Skip() //remove for activating it
	queues := []string{constants.REQUEST, constants.REPLY}
	checkMode(queues)
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	rabbitConn, err := connection.New(constants.AMQPURL)
	rabbitConn.PurgeAll()
	assert.Equal(err, nil, "Shouldn't be an error")
	chBatch, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo, err := spaceRepo.New(rabbitConn)
	msgs, _ := chBatch.Consume(
		queues[0], // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	corrId := "-1"
	go func() {
		for resp := range msgs {
			corrId = resp.CorrelationId
			response, _ := json.Marshal("1")
			chBatch.Publish(
				"",              // exchange
				queues[1], // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
			//resp.Ack(false)
			break
		}
	}()

	done, err := spaceRepo.ReserveBatch([]string{}, domain.Hour{Hour: 12, Min: 30}, domain.Hour{Hour: 13, Min: 30}, []string{s}, "Iñigol")
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.NotEqual(done, "-1", "Should be positive")
}

func TestFilterBy(t *testing.T) {
	t.Skip() //remove for activating itç
	queues := []string{constants.REQUEST, constants.REPLY}
	checkMode(queues)
	assert := assert.New(t)
	rabbitConn, err := connection.New(constants.AMQPURL)
	rabbitConn.PurgeAll()
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo, _ := spaceRepo.New(rabbitConn)
	msgs, _ := chReserve.Consume(
		queues[0], // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	messageSent := []domain.Space{
		{
			Name:     "A1",
			Capacity: 20,
			Building: "Ada",
			Kind:     "aula",
		},}
	corrId := "-1"
	go func() {
		for resp := range msgs {
			corrId = resp.CorrelationId
			response, _ := json.Marshal(messageSent)
			chReserve.Publish(
				"",              // exchange
				queues[1], // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
			//resp.Ack(false)
			break
		}
	}()

	messageRecieved, err := spaceRepo.FilterBy(domain.SpaceFilterParams{Capacity: 5, Day: "2022-01-02", Hour: domain.Hour{Hour: 12, Min: 0},Floor: "1", Building: "Ada"})
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(messageRecieved, messageSent, "Should be true")
}