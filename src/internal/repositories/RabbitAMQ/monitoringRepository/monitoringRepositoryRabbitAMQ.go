package monitoringrepositoryrabbitamq

import (
	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type MonitoringRepositorioRabbit struct {
	*rabbitamqRepository.Repository
}

func New(ch *amqp.Channel) *MonitoringRepositorioRabbit {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(ch, queues)
	if err != nil {
		return nil
	}
	return &MonitoringRepositorioRabbit{rp}
}

func (repo *MonitoringRepositorioRabbit) Ping() (bool, error) {
	_, err := repo.RCPcallJSON("", constants.PING)
	if err != nil {
		return false, err
	}

	return true, nil
}
