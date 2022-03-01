package rabbitMQ

import (
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/streadway/amqp"
)

type HorarioRepositorioRabbit struct {
	conn *amqp.Connection
	ch *amqp.Channel
}

func New(address string) *HorarioRepositorioRabbit {
	conn, ch, _ := connection.Connect(address)
	return &HorarioRepositorioRabbit{conn, ch}
}

func (repo *HorarioRepositorioRabbit) CloseConn() {
	connection.Disconnect(repo.conn, repo.ch)
}

func (repo *HorarioRepositorioRabbit) Monitoring(){
	
}