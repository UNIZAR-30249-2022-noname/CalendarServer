package issuerepositoryrabbitamq

import (
	"encoding/json"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	rabbitamqRepository "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/RabbitAMQ"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/streadway/amqp"
)

type IssueRepository struct {
	*rabbitamqRepository.Repository
}

func New(ch *amqp.Channel) (*IssueRepository, error) {
	queues := []string{constants.REQUEST, constants.REPLY}
	rp, err := rabbitamqRepository.New(ch, queues)
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
	responseJSON, err := repo.RCPcallJSON(key, constants.DELETEISSUE)
	if err != nil {
		return err
	}
	var reply rabbitamqRepository.DataMessageQueue[string]
	json.Unmarshal(responseJSON, &reply)
	if reply.Response.Result != "ok" {
		return apperrors.ErrNotFound
	}
	return nil
}

func (repo *IssueRepository) Create(issue domain.Issue) error {
	responseJSON, err := repo.RCPcallJSON(issue, constants.CREATE_ISSUE)
	if err != nil {
		return err
	}
	var reply rabbitamqRepository.DataMessageQueue[string]
	json.Unmarshal(responseJSON, &reply)
	if reply.Response.Result != "ok" {
		return apperrors.ErrNotFound
	}
	return nil
}
func (repo *IssueRepository) ChangeState(key string, state int) error {
	type issueUpdateType struct {
		Key   string `json:"key"`
		State int    `json:"state"`
	}
	updateIssue := issueUpdateType{Key: key, State: state}
	responseJSON, err := repo.RCPcallJSON(updateIssue, constants.UPDATEISSUE)
	if err != nil {
		return err
	}
	var reply rabbitamqRepository.DataMessageQueue[string]
	json.Unmarshal(responseJSON, &reply)
	if reply.Response.Result != "ok" {
		return apperrors.ErrNotFound
	}
	return nil
}
