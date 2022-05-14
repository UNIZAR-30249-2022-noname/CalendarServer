package issue

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
)

type IssueServiceImp struct {
	issueRepository ports.IssueRepository
}

func New(issueRepository ports.IssueRepository) *IssueServiceImp {
	return &IssueServiceImp{issueRepository: issueRepository}
}

func (svc *IssueServiceImp) GetAll() ([]domain.Issue, error) {

	return svc.issueRepository.GetAll()

}
func (svc *IssueServiceImp) Delete(key string) error {

	return svc.issueRepository.Delete(key)

}

func (svc *IssueServiceImp) Create(issue domain.Issue) error {

	return svc.issueRepository.Create(issue)

}

func (svc *IssueServiceImp) ChangeState(key string, state int) error {

	return svc.issueRepository.ChangeState(key, state)

}

func (svc *IssueServiceImp) DownloadIssues() ([]byte ,error) {

	return svc.issueRepository.DownloadIssues()

}