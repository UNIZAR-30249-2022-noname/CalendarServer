package issuerepositorymemory

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type IssueRepository struct {
}

func New() *IssueRepository {
	return &IssueRepository{}
}

func (repo *IssueRepository) GetAll() ([]domain.Issue, error) {
	return []domain.Issue{
		{
			Tags:        []string{"Urgente"},
			Title:       "goteras",
			Description: "Cae agua del techo",
			Key:         "1",
			Space:       "A0.11",
			State:       0,
		},
		{
			Tags:        []string{"Urgente"},
			Title:       "goteras",
			Description: "Cae agua del techo",
			Key:         "2",
			Space:       "A0.11",
			State:       0,
		},
		{
			Tags:        []string{"Urgente"},
			Title:       "Enchufe",
			Description: "Cae agua del techo",
			Key:         "3",
			Space:       "A0.11",
			State:       1,
		},
		{
			Tags:        []string{"Urgente"},
			Title:       "proyector",
			Description: "Cae agua del techo",
			Key:         "4",
			Space:       "A0.11",
			State:       2,
		},
		{
			Tags:        []string{"Urgente"},
			Title:       "proyector",
			Description: "Cae agua del techo",
			Key:         "5",
			Space:       "A0.11",
			State:       2,
		},
	}, nil
}

func (repo *IssueRepository) Delete(key string) error {
	return nil
}

func (repo *IssueRepository) Create(issue domain.Issue) error {
	return nil
}
func (repo *IssueRepository) ChangeState(key string, state int) error {
	return nil
}
