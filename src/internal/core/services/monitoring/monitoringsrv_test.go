package monitoring_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/monitoring"
	horariorepositoriorabbit "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio/rabbitMQ/repoRabbit"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/stretchr/testify/assert"
)


func TestPing(t *testing.T) {
	var rabbitConn connection.Connection
	rabbitConn, err := connection.New(constants.AMQPURL)
	chScheduler, err := rabbitConn.NewChannel()
	defer rabbitConn.Disconnect()
	horariorepoRMQ := horariorepositoriorabbit.New(chScheduler)
	rabbit := monitoring.New(horariorepoRMQ)
	res, err := rabbit.Monitoring()
	assert.Equal(t, res, true)
	assert.Equal(t, err, nil)
}