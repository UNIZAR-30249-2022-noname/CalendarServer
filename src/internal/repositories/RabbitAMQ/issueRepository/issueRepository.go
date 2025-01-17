package issuerepositoryrabbitamq

import (
	"encoding/json"
	"fmt"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	connection "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/connect"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
)

type IssueRepository struct {
	*rabbitamqRepository.Repository
}

func New(rabbitConn connection.Connection) (*IssueRepository, error) {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(rabbitConn, queues)
	if err != nil {
		return &IssueRepository{}, err
	}
	return &IssueRepository{rp}, nil
}

func (repo *IssueRepository) GetAll() ([]domain.Issue, error) {
	var reply rabbitamqRepository.DataMessageQueue[[]domain.Issue]
	allIssuesJSON, err := repo.RCPcallJSON("", constants.GETISSUES)
	if err != nil {
		return []domain.Issue{}, err
	}
	json.Unmarshal(allIssuesJSON, &reply)
	return reply.Response.Result, nil
}

func (repo *IssueRepository) Delete(key string) error {
	type issueDeleteType struct {
		Key string `json:"key"`
	}
	issue := issueDeleteType{Key: key}
	_, err := repo.RCPcallJSON(issue, constants.DELETEISSUE)
	if err != nil {
		return err
	}

	return nil
}

func (repo *IssueRepository) Create(issue domain.Issue) error {
	responseJSON, err := repo.RCPcallJSON(issue, constants.NEWISSUE)
	if err != nil {
		return err
	}
	var reply rabbitamqRepository.DataMessageQueue[string]
	json.Unmarshal(responseJSON, &reply)
	return nil
}
func (repo *IssueRepository) ChangeState(key string, state int) error {
	type issueUpdateType struct {
		Key   string `json:"key"`
		State int    `json:"state"`
	}
	updateIssue := issueUpdateType{Key: key, State: state}
	_, err := repo.RCPcallJSON(updateIssue, constants.UPDATEISSUE)
	if err != nil {
		return err
	}

	return nil
}

func (repo *IssueRepository) DownloadIssues(building string) ([]byte, error) {
	var reply rabbitamqRepository.DataMessageQueueDownload
	allIssuesJSON, err := repo.RCPcallJSON(building, constants.DOWNLOADISSUE)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Tipo: %T\n", allIssuesJSON)
	json.Unmarshal(allIssuesJSON, &reply)
	return reply.Response.Result.Data, nil
}
