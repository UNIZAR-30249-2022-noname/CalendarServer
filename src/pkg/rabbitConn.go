package pkg

import "github.com/streadway/amqp"

func connect(address string) (*amqp.Connection, *amqp.Channel, error) {
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

func disconnect(conn *amqp.Connection, ch *amqp.Channel) {
    conn.Close()
    ch.Close()
}