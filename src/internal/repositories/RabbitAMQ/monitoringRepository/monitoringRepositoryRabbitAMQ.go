package monitoringrepositoryrabbitamq

import (
	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
)

type MonitoringRepositorioRabbit struct {
	*rabbitamqRepository.Repository
}

func New(rabbitConn connection.Connection) *MonitoringRepositorioRabbit {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(rabbitConn, queues)
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
