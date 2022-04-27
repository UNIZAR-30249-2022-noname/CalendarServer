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
	var availableHours = []domain.AvailableHours{}
	availableHoursJSON, err := repo.RCPcallJSON(req, constants.GETAVAILABLEHOURS)
	if err != nil {
		return []domain.AvailableHours{}, err
	}
	json.Unmarshal(availableHoursJSON, &availableHours)
	return availableHours, nil

}
func (repo *SchedulerRepository) UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error) {
	req := struct {
		terna   domain.DegreeSet
		entries []domain.Entry
	}{
		entries: entries,
		terna:   terna,
	}

	reply := struct {
		id string
	}{}
	availableHoursJSON, err := repo.RCPcallJSON(req, constants.UPDATESCHEDULER)
	if err != nil {
		return "", err
	}
	json.Unmarshal(availableHoursJSON, &reply)
	return reply.id, nil

}
func (repo *SchedulerRepository) DeleteEntry(domain.Entry) error                      {}
func (repo *SchedulerRepository) ListAllDegrees() ([]domain.DegreeDescription, error) {}
func (repo *SchedulerRepository) DeleteAllEntries(terna domain.DegreeSet) error       {}
func (repo *SchedulerRepository) GetEntries(domain.DegreeSet) ([]domain.Entry, error) {}
