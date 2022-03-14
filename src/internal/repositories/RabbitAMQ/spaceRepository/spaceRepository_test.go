package spaceRepository_test

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

func TestReserve(t *testing.T) {
	//t.Skip() //remove for activating it
	assert := assert.New(t)
	
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chReserve, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo := spaceRepo.New(chReserve)
	done, err := spaceRepo.Reserve(domain.Space{},domain.Hour{Hour: 12, Min: 30},domain.Hour{Hour: 13, Min: 30})
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(done, true, "Should be true")
	chReserve.QueueDelete(constants.RESERVE,true,false,true)
}

func TestReserveBatch(t *testing.T) {
	//t.Skip() //remove for activating it
	assert := assert.New(t)
	a := time.Now().Local()
	s := a.Format("2006-01-02")
	rabbitConn, err := connection.New(constants.AMQPURL)
	assert.Equal(err, nil, "Shouldn't be an error")
	chBatch, err := rabbitConn.NewChannel()
	assert.Equal(err, nil, "Shouldn't be an error")
	err = connection.PrepareChannel(chBatch, constants.BATCH)
	assert.Equal(err, nil, "Shouldn't be an error")
	err = connection.PrepareChannel(chBatch, constants.BATCH_REPLY)
	assert.Equal(err, nil, "Shouldn't be an error")
	spaceRepo := spaceRepo.New(chBatch)
	msgs, _ := chBatch.Consume(
		constants.BATCH_REPLY, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
		corrId := "-1"
		go func (){
			for resp := range msgs{
				corrId = resp.CorrelationId
				response, _ := json.Marshal("1")
				chBatch.Publish(
				"",          // exchange
				constants.BATCH_REPLY, // routing key
				false,       // mandatory
				false,       // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: corrId,
					Body:          response,
				})
				resp.Ack(false)
			}
		}()
		
	done, err := spaceRepo.ReserveBatch([]domain.Space{},domain.Hour{Hour: 12, Min: 30},domain.Hour{Hour: 13, Min: 30},[]string{s})
	assert.Equal(err, nil, "Shouldn't be an error")
	assert.Equal(done, true, "Should be true")
	chBatch.QueueDelete(constants.BATCH,true,false,true)
}