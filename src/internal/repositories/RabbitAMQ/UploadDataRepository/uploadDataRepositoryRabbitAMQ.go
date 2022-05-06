package uploaddatarepositoryrabbitamq

import (
	"encoding/json"

	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
)

type UploadDataRepository struct {
	*rabbitamqRepository.Repository
}

func New(rabbitConn connection.Connection) (*UploadDataRepository, error) {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(rabbitConn, queues)
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