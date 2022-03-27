package spaceRepository

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/auxFuncs"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type SpaceRepository struct {
	ch *amqp.Channel
}

func New(ch *amqp.Channel) *SpaceRepository {
	return &SpaceRepository{ch}
}

func (repo *SpaceRepository) RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error) {
	err := connect.PrepareChannel(repo.ch, constants.REQ_INFO_SLOT)
	if err != nil {
		return domain.AllInfoSlot{}, err
	}
	msg, err := json.Marshal(req)
	if err != nil {
		return domain.AllInfoSlot{}, err
	}
	err = repo.ch.Publish(
		"",          // exchange
		constants.REQ_INFO_SLOT, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: auxFuncs.RandomString(10),
			Body:          msg,
		})

		sd := domain.SlotData{
			Name: "A1",
			Capacity: 5,
			Description: "Lorem ipsum no leas mas porque esto es dummy text",
			Building: "Ada",
			Floor: "baja",
			Type: "aula",
		  };

		  
		  is := []domain.InfoSlots{
			{
			Hour: 8,
			Occupied: false,
			},
			{
				Hour: 9,
				Occupied: true,
				Person: "Urrikote",
			},
			{
				Hour: 10,
				Occupied: false,
			},
			{
				Hour: 11,
				Occupied: false,
			},
			{
				Hour: 12,
				Occupied: true,
				Person: "Urrikyu",
			},
			{
				Hour: 13,
				Occupied: false,
			},
			{
				Hour: 14,
				Occupied: false,
			},
			{
				Hour: 15,
				Occupied: true,
				Person: "Urriuuuu",
			},
			{
				Hour: 16,
				Occupied: false,
			},
			{
				Hour: 17,
				Occupied: false,
			},
			{
				Hour: 8,
				Occupied: true,
				Person: "Urrikoncio",
			},
			{
				Hour: 19,
				Occupied: false,
			},
			{
				Hour: 20,
				Occupied: false,
			},
		}
	
		allInfo := domain.AllInfoSlot{
			SlotData: sd,
			InfoSlots: is,
		}

	return allInfo, err
}

func (repo *SpaceRepository) Reserve(sp domain.Space, init domain.Hour, date, person string) (string, error) {
	err := connect.PrepareChannel(repo.ch, constants.RESERVE)
	if err != nil {
		return "-1", err
	}
	msg, err := json.Marshal(domain.Reserve{Space: sp, Init: init, Date: date, Person: person})
	if err != nil {
		return "-1", err
	}
	err = repo.ch.Publish(
		"",          // exchange
		constants.RESERVE, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: auxFuncs.RandomString(10),
			Body:          msg,
		})
	return "1", err
}


func (repo *SpaceRepository) ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string, person string) (string, error) {

	msg, err := json.Marshal(domain.ReserveBatch{Spaces: spaces, Init: init, End: end, Dates: dates, Person: person})
	if err != nil {
		return "-1", err
	}

	msgs, err := repo.ch.Consume(
		constants.BATCH_REPLY, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	corrId := auxFuncs.RandomString(10)
	err = repo.ch.Publish(
		"",          // exchange
		constants.BATCH, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo: constants.BATCH_REPLY,
			Body:          msg,
		})
	if err != nil {
		return "-1", err
	}
	
	
	lastId := "-1"
	for resp := range msgs {
		if corrId == resp.CorrelationId {
			if err == nil {
				json.Unmarshal(resp.Body, &lastId)
			}
			break
		}
	}
	if lastId == "-1" {
		return "-1", err
	}
	return lastId, err
}
