package connect

import "github.com/streadway/amqp"

func Connect(address string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, nil, err
	}
    ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	return conn, ch, err
}

func Disconnect(conn *amqp.Connection, ch *amqp.Channel) {
    conn.Close()
    ch.Close()
}