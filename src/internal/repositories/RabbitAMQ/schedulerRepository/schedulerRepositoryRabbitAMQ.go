package schedulerrepositoryrabbitamq

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
)

type SchedulerRepository struct {
	*rabbitamqRepository.Repository
}

func New(rabbitConn connection.Connection) (*SchedulerRepository, error) {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(rabbitConn, queues)
	if err != nil {
		return &SchedulerRepository{}, err
	}
	return &SchedulerRepository{rp}, nil
}

type DegreeSetAux struct {
	DegreeSet domain.DegreeSet
}

func (repo *SchedulerRepository) GetAvailableHours(req domain.DegreeSet) ([]domain.AvailableHours, error) {
	var reply rabbitamqRepository.DataMessageQueue[[]domain.AvailableHours]
	availableHoursJSON, err := repo.RCPcallJSON(DegreeSetAux{req}, constants.GETAVAILABLEHOURS)
	if err != nil {
		return []domain.AvailableHours{}, err
	}
	json.Unmarshal(availableHoursJSON, &reply)
	return reply.Response.Result, nil

}
func (repo *SchedulerRepository) UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error) {
	req := struct {
		DegreeSet   domain.DegreeSet
		Entry []domain.Entry
	}{
		Entry: entries,
		DegreeSet:  terna,
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
	availableHoursJSON, err := repo.RCPcallJSON(DegreeSetAux{req}, constants.GETENTRIES)
	if err != nil {
		return []domain.Entry{}, err
	}
	json.Unmarshal(availableHoursJSON, &reply)
	return reply.Response.Result, nil

}

func (repo *SchedulerRepository) GetICS(terna domain.DegreeSet) (string, error) {
	var serializedIcal string
	res, err := repo.RCPcallJSON(terna, constants.GETICS)
	json.Unmarshal(res, &serializedIcal)
	return serializedIcal, err;
}
