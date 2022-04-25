package connect

import (
	"os"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type Connection struct {
	connection *amqp.Connection
	channels   []*amqp.Channel
	open       bool
}

func New(address string) (Connection, error) {
	conn, err := amqp.Dial(address)
	if err != nil {
		return Connection{}, err //TODO poner un error nuestro
	}
	return Connection{connection: conn, channels: []*amqp.Channel{}, open: true}, err
}
func (conn *Connection) Disconnect() {
	for _, ch := range conn.channels {
		ch.Close()
	}
	conn.connection.Close()
	conn.open = false

}

func (conn *Connection) NewChannel() (*amqp.Channel, error) {
	ch, err := conn.connection.Channel()
	if err != nil {
		return nil, err //TOSO poner un error nuestro
	}
	conn.channels = append(conn.channels, ch)
	return ch, err

}

func PrepareChannel(ch *amqp.Channel, qname string) (error){
	_, err := ch.QueueDeclare(
		qname, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return apperrors.ErrConn
	}
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return apperrors.ErrConn
	}
	return nil
}

func (conn *Connection) PurgeAll() (error) {
	ch, _ := conn.NewChannel()
	queues := []string{constants.REQUEST, constants.REPLY}
	if os.Getenv("GATEWAY_MODE") == constants.TEST_MODE {
		for i := 0; i < len(queues); i++ {
			queues[i] += "_test"
		}
	}
	nrep, error := ch.QueuePurge(queues[0], true)
	if nrep < 0 {
		return error
	}
	_, error = ch.QueuePurge(queues[1], true)
	return error
}
