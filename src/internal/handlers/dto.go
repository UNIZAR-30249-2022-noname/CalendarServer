package handlers

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type TernaDto struct {
	Degree string `json:"degree"`
	Year   int    `json:"year"`
	Group  string `json:"group"`
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
		Week:    e.Week,
		Group:   e.Group,
	}

}

func EntryDomainToDTO(d domain.Entry) EntryDTO {
	return EntryDTO{
		InitHour: d.Init.Hour,
		InitMin:  d.Init.Min,
		EndHour:  d.End.Hour,
		EndMin:   d.End.Min,
		Subject:  d.Subject.Name,
		Kind:     d.Subject.Kind,
		Room:     d.Room.Name,
		Week:     d.Week,
		Group:    d.Group,
		Weekday:  d.Weekday,
	}

}

func EntriesDTOtoDomain(dtos []EntryDTO) []domain.Entry {
	ls := []domain.Entry{}
	for _, e := range dtos {
		ls = append(ls, e.ToEntry())
	}
	return ls

}

func EntriesDomaintoDTO(entries []domain.Entry) []EntryDTO {
	ls := []EntryDTO{}
	for _, e := range entries {
		ls = append(ls, EntryDomainToDTO(e))
	}
	return ls

}

type ListDegreesDTO struct {
	List []domain.DegreeDescription `json:"list"`
}

func NewListDegrees(l []domain.DegreeDescription) ListDegreesDTO {
	return ListDegreesDTO{List: l}
}

type ErrorHttp struct {
	Message string `json:"message"`
}
