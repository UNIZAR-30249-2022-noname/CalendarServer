package repoRabbit

import (
	"github.com/streadway/amqp"
)

type HorarioRepositorioRabbit struct {
	ch *amqp.Channel
}

func New(ch *amqp.Channel) *HorarioRepositorioRabbit {
	return &HorarioRepositorioRabbit{ch}
}

func (repo *HorarioRepositorioRabbit) CloseConn() {
}

func (repo *HorarioRepositorioRabbit) Monitoring() (bool, error){
	return true, nil
}