package schedulerrepositoryrabbitamq

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type SchedulerRepository struct {
	*rabbitamqRepository.Repository
}

func New(ch *amqp.Channel) (*SchedulerRepository, error) {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(ch, queues)
	if err != nil {
		return &SchedulerRepository{}, err
	}
	return &SchedulerRepository{rp}, nil
}

func (repo *SchedulerRepository) GetAvailableHours(req domain.DegreeSet) ([]domain.AvailableHours, error) {
	var reply rabbitamqRepository.DataMessageQueue[[]domain.AvailableHours]
	availableHoursJSON, err := repo.RCPcallJSON(req, constants.GETAVAILABLEHOURS)
	if err != nil {
		return []domain.AvailableHours{}, err
	}
	json.Unmarshal(availableHoursJSON, &reply)
	return reply.Response.Result, nil

}
func (repo *SchedulerRepository) UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error) {
	req := struct {
		terna   domain.DegreeSet
		entries []domain.Entry
	}{
		entries: entries,
		terna:   terna,
	}

	var reply rabbitamqRepository.DataMessageQueue[string]
	availableHoursJSON, err := repo.RCPcallJSON(req, constants.UPDATESCHEDULER)
	if err != nil {
		return "", err
	}
	json.Unmarshal(availableHoursJSON, &reply)
	return reply.Response.Result, nil

}
func (repo *SchedulerRepository) DeleteEntry(req domain.Entry) error {
	_, err := repo.RCPcallJSON(req, constants.DELETEENTRY)
	return err

}
func (repo *SchedulerRepository) ListAllDegrees() ([]domain.DegreeDescription, error) {
	var reply rabbitamqRepository.DataMessageQueue[[]domain.DegreeDescription]
	availableHoursJSON, err := repo.RCPcallJSON(nil, constants.LISTALLDEGREES)
	if err != nil {
		return []domain.DegreeDescription{}, err
	}
	json.Unmarshal(availableHoursJSON, &reply)
	return reply.Response.Result, nil

}
func (repo *SchedulerRepository) DeleteAllEntries(terna domain.DegreeSet) error {
	_, err := repo.RCPcallJSON(terna, constants.DELETEALLENTRIES)
	return err
}
func (repo *SchedulerRepository) GetEntries(req domain.DegreeSet) ([]domain.Entry, error) {

	var reply rabbitamqRepository.DataMessageQueue[[]domain.Entry]
	availableHoursJSON, err := repo.RCPcallJSON(req, constants.GETENTRIES)
	if err != nil {
		return []domain.Entry{}, err
	}
	json.Unmarshal(availableHoursJSON, &reply)
	return reply.Response.Result, nil

}
