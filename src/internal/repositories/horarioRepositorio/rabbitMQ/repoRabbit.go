package rabbitMQ

import (
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

type HorarioRepositorioMySQL struct {
	conn *amqp.Connection
	ch *amqp.Channel
}

func New(address string) *HorarioRepositorioMySQL {
	conn, ch, _ := connection.Connect(address)
	return &HorarioRepositorioMySQL{conn, ch}
}

func (repo *HorarioRepositorioMySQL) CloseConn() {
	connection.Disconnect(repo.conn, repo.ch)
}

