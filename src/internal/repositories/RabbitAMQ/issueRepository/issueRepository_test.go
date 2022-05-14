package issuerepositoryrabbitamq_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	issueRepo "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ/issueRepository"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestDeleteQueueBeforeTest(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReqInfo, err := rabbitConn.NewChannel()
	chReqInfo.QueueDelete(constants.REQUEST, true, false, true)
	chReqInfo.QueueDelete(constants.REPLY, true, false, true)
}

func TestGetAllIssues(t *testing.T) {
	t.Skip() //remove for activating it
	assert := assert.New(t)
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	issueRepo, _ := issueRepo.New(rabbitConn)
	msgs, _ := chReserve.Consume(
		constants.REQUEST, // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	messageSent := []domain.Issue{
		{
			Tags:        []string{"Urgente"},
			Title:       "goteras",
			Description: "Cae agua del techo",
			Key:         "1",
			Space:       "A0.11",
			State:       0,
		},
		{
			Tags:        []string{"Urgente"},
			Title:       "goteras",
			Description: "Cae agua del techo",
			Key:         "2",
			Space:       "A0.11",
			State:       0,
		}}
	corrId := "-1"
	go func() {
		for resp := range msgs {
			corrId = resp.CorrelationId
			response, _ := json.Marshal(messageSent)
			chReserve.Publish(
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
			break
		}
	}()

	messageRecieved, err := issueRepo.GetAll()
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(messageRecieved, messageSent, "Should be true")
	chReserve.QueueDelete(constants.REQUEST, true, false, true)
	chReserve.QueueDelete(constants.REPLY, true, false, true)
}

func TestDeleteIssue(t *testing.T) {
	t.Skip() //remove for activating it
	assert := assert.New(t)
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	issueRepo, _ := issueRepo.New(rabbitConn)
	msgs, _ := chReserve.Consume(
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
		i := 0
		for resp := range msgs {
			okno := "ok"
			if i != 0 {
				okno = "nook"
			}
			corrId = resp.CorrelationId
			response, _ := json.Marshal(okno)
			chReserve.Publish(
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
			i++
			if i >= 2 {
				break
			}
		}
	}()

	err = issueRepo.Delete("OWO")
	assert.Equal(err, nil, "Shouldn't be an error")
	err = issueRepo.Delete("NOWO")
	assert.Equal(err, apperrors.ErrNotFound, "Shouldn't be an error")
	chReserve.QueueDelete(constants.REQUEST, true, false, true)
	chReserve.QueueDelete(constants.REPLY, true, false, true)
}

func TestCreateIssue(t *testing.T) {
	t.Skip() //remove for activating it
	assert := assert.New(t)
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	issueRepo, _ := issueRepo.New(rabbitConn)
	msgs, _ := chReserve.Consume(
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
		i := 0
		for resp := range msgs {
			okno := "ok"
			if i != 0 {
				okno = "nook"
			}
			corrId = resp.CorrelationId
			response, _ := json.Marshal(okno)
			chReserve.Publish(
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
			i++
			if i >= 2 {
				break
			}
		}
	}()

	issue := domain.Issue{

		Tags:        []string{"Urgente"},
		Title:       "goteras",
		Description: "Cae agua del techo",
		Key:         "1",
		Space:       "A0.11",
		State:       0,
	}

	err = issueRepo.Create(issue)
	assert.Equal(err, nil, "Shouldn't be an error")
	err = issueRepo.Create(issue)
	assert.Equal(err, apperrors.ErrNotFound, "Shouldn't be an error")
	chReserve.QueueDelete(constants.REQUEST, true, false, true)
	chReserve.QueueDelete(constants.REPLY, true, false, true)
}

func TestChangeState(t *testing.T) {
	t.Skip() //remove for activating it
	assert := assert.New(t)
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	issueRepo, _ := issueRepo.New(rabbitConn)
	msgs, _ := chReserve.Consume(
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
		i := 0
		for resp := range msgs {
			okno := "ok"
			if i != 0 {
				okno = "nook"
			}
			corrId = resp.CorrelationId
			response, _ := json.Marshal(okno)
			chReserve.Publish(
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
			i++
			if i >= 2 {
				break
			}
		}
	}()

	err = issueRepo.ChangeState("1", 1)
	assert.Equal(err, nil, "Shouldn't be an error")
	err = issueRepo.ChangeState("2", 2)
	assert.Equal(err, apperrors.ErrNotFound, "Shouldn't be an error")
	chReserve.QueueDelete(constants.REQUEST, true, false, true)
	chReserve.QueueDelete(constants.REPLY, true, false, true)
}

func TestDownloadIssues(t *testing.T) {
	t.Skip() //remove for activating it
	assert := assert.New(t)
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	err = connection.PrepareChannel(chReserve, constants.REQUEST)
	assert.Equal(err, nil, "Shouldn't be an error")
	err = connection.PrepareChannel(chReserve, constants.REPLY)
	assert.Equal(err, nil, "Shouldn't be an error")
	issueRepo, _ := issueRepo.New(chReserve)
	msgs, _ := chReserve.Consume(
		constants.REQUEST, // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	messageSent, _ := ioutil.ReadFile("C:/Users/Equipo/Desktop/LIS/Gateway/src/internal/handlers/prueba/file.pdf")
	corrId := "-1"
	go func() {
		for resp := range msgs {
			corrId = resp.CorrelationId
			response, _ := json.Marshal(messageSent)
			chReserve.Publish(
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
			break
		}
	}()

	messageRecieved, err := issueRepo.DownloadIssues()
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(messageRecieved, messageSent, "Should be true")
	chReserve.QueueDelete(constants.REQUEST, true, false, true)
	chReserve.QueueDelete(constants.REPLY, true, false, true)
}
