package connect

import (
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
