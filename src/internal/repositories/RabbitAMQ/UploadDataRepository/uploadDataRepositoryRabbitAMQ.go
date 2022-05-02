package uploaddatarepositoryrabbitamq

import (
	"encoding/json"

	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type UploadDataRepository struct {
	*rabbitamqRepository.Repository
}

func New(ch *amqp.Channel) (*UploadDataRepository, error) {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(ch, queues)
	if err != nil {
		return &UploadDataRepository{}, err
	}
	return &UploadDataRepository{rp}, nil
}

func (repo *UploadDataRepository) UpdateByCSV(req string) (bool, error) {
	responseJSON, err := repo.RCPcallJSON(req, constants.IMPORT)
	if err != nil {
		return false, err
	}
	response := false
	json.Unmarshal(responseJSON, &response)
	return response, nil

}