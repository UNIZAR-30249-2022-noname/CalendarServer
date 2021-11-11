package handlers

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type TernaDto struct {
	Titulacion string `json:"titulacion"`
	Curso      int    `json:"curso"`
	Grupo      int    `json:"grupo"`
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
}

func (e EntryDTO) ToEntry() domain.Entry {
	return domain.Entry{
		Init:    domain.NewHour(e.InitHour, e.InitMin),
		End:     domain.NewHour(e.EndHour, e.EndMin),
		Subject: domain.Subject{Kind: e.Kind, Name: e.Subject},
		Room:    domain.Room{Name: e.Room},
	}

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
