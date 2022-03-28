package spacerepositoryrabbitamq_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	spaceRepo "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ/spaceRepository"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestRequestInfoSlots(t *testing.T) {
	//t.Skip() //remove for activating it
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	data := domain.ReqInfoSlot{Name: "A1", Date: s}
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReqInfo, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo, err := spaceRepo.New(chReqInfo)

	//Simulated server
	msgs, _ := chReqInfo.Consume(
		constants.REQUEST, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	myResponse := domain.AllInfoSlot{
		SlotData: domain.SlotData{
			Name:        "A1",
			Capacity:    5,
			Description: "Lorem ipsum no leas mas porque esto es dummy text",
			Building:    "Ada",
			Floor:       "baja",
			Type:        "aula",
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
				constants.REPLY, // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
			//resp.Ack(false)
		}
	}()

	//RUN
	done, err := spaceRepo.RequestInfoSlots(data)
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(done, myResponse, "Should be positive")
	chReqInfo.QueueDelete(constants.REQUEST, true, false, true)
	chReqInfo.QueueDelete(constants.REPLY, true, false, true)
}

func TestRequestInfoSlotsMultiple(t *testing.T) {
	//t.Skip() //remove for activating it
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	data := domain.ReqInfoSlot{Name: "A1", Date: s}
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReqInfo, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo, err := spaceRepo.New(chReqInfo)

	//Simulated server
	msgs, _ := chReqInfo.Consume(
		constants.REQUEST, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	myResponse := domain.AllInfoSlot{
		SlotData: domain.SlotData{
			Name:        "A1",
			Capacity:    5,
			Description: "Lorem ipsum no leas mas porque esto es dummy text",
			Building:    "Ada",
			Floor:       "baja",
			Type:        "aula",
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
				constants.REPLY, // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
			//resp.Ack(false)
		}
	}()

	//RUN
	done, err := spaceRepo.RequestInfoSlots(data)
	done, err = spaceRepo.RequestInfoSlots(data)
	done, err = spaceRepo.RequestInfoSlots(data)
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(done, myResponse, "Should be positive")
	chReqInfo.QueueDelete(constants.REQUEST, true, false, true)
	chReqInfo.QueueDelete(constants.REPLY, true, false, true)
}
func TestReserve(t *testing.T) {
	//t.Skip() //remove for activating it
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo, _ := spaceRepo.New(chReserve)
	done, err := spaceRepo.Reserve(domain.Space{}, domain.Hour{Hour: 12, Min: 30}, domain.Hour{Hour: 12, Min: 30}, s, "Iñigol")
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(done, "1", "Should be true")
	//chReserve.QueueDelete(constants.RESERVE,true,false,true)
}

func TestReserveBatch(t *testing.T) {
	t.Skip() //remove for activating it
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chBatch, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	err = connection.PrepareChannel(chBatch, constants.REQUEST)
	assert.Equal(err, nil, "Shouldn't be an error")
	err = connection.PrepareChannel(chBatch, constants.REPLY)
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo, err := spaceRepo.New(chBatch)
	msgs, _ := chBatch.Consume(
		constants.REQUEST, // queue
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
				constants.REPLY, // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
			resp.Ack(false)
		}
	}()

	done, err := spaceRepo.ReserveBatch([]domain.Space{}, domain.Hour{Hour: 12, Min: 30}, domain.Hour{Hour: 13, Min: 30}, []string{s}, "Iñigol")
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.NotEqual(done, "-1", "Should be positive")
	chBatch.QueueDelete(constants.REQUEST, true, false, true)
	chBatch.QueueDelete(constants.REPLY, true, false, true)
}
