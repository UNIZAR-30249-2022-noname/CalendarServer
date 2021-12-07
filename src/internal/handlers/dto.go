package handlers

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type TernaDto struct {
	Titulacion string `json:"titulacion"`
	Curso      int    `json:"curso"`
	Grupo      string `json:"grupo"`
}

type SchedulerDTO struct {
	AvailableHours []domain.AvailableHours `json:"availableHours"`
}

func NewScheduler(availableHours []domain.AvailableHours) SchedulerDTO {
	return SchedulerDTO{AvailableHours: availableHours}

}

type EntryDTO struct {
	InitHour int    `json:"initHour"`
	InitMin  int    `json:"initMin"`
	EndHour  int    `json:"endHour"`
	EndMin   int    `json:"endMin"`
	Subject  string `json:"subject"`
	Kind     int    `json:"kind"`
	Room     string `json:"room"`
	Week     string `json:"semana"`
	Group    string `json:"grupo"`
	Weekday  int    `json:"weekday"`
}

func (e EntryDTO) ToEntry() domain.Entry {
	return domain.Entry{
		Init:    domain.NewHour(e.InitHour, e.InitMin),
		End:     domain.NewHour(e.EndHour, e.EndMin),
		Subject: domain.Subject{Kind: e.Kind, Name: e.Subject},
		Room:    domain.Room{Name: e.Room},
		Weekday: e.Weekday,
	}

}

func EntriesDTOtoDomain(dtos []EntryDTO) []domain.Entry {
	ls := []domain.Entry{}
	for _, e := range dtos {
		ls = append(ls, e.ToEntry())
	}
	return ls

}

type ErrorHttp struct {
	Message string `json:"message"`
}
