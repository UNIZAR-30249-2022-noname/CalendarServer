package domain

import "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"

const (
	THEORICAL = 1
	PRACTICES = 2
	EXERCISES = 3
)

//AvaialableHours is a struct which represents the available hours
//per [Terna]
type AvailableHours struct {
	Subject   Subject
	Remaining int
	Max       int
}
type Subject struct {
	Kind int
	Name string
}

func (s Subject) IsValid() error {
	if s.Kind == 0 || s.Name == "" {
		return apperrors.ErrInvalidInput
	}
	return nil
}

//Terna is a struct which represent the relation among
// bachelors, year and group
type Terna struct {
	Titulacion string
	Curso      int
	Grupo      string
}

type Hour struct {
	hour int
	min  int
}

func NewHour(h, m int) Hour {
	return Hour{
		hour: h,
		min:  m,
	}
}

func (h Hour) IsLaterThan(h2 Hour) bool {
	//if the hour is previus return false
	if h.hour < h2.hour {
		return false
		//if the hour is equal check the minutes
	} else if h.hour == h2.hour && h.min <= h2.min {
		return false
	}
	return true
}
func HourToInt(h Hour) int {
	return h.hour*100 + h.min
}

type Room struct {
	Name string
}

type Entry struct {
	Init    Hour
	End     Hour
	Subject Subject
	Room    Room
	Week    string
	Group   string
}

func (e Entry) IsValid() error {

	//check if there is not empty compulsory fields
	if (e.Init == Hour{}) || (e.End == Hour{}) || (e.Subject == Subject{}) {
		return apperrors.ErrInvalidInput
	}
	//Check if the entry has valid time interval
	if e.Init.IsLaterThan(e.End) {
		return apperrors.ErrInvalidInput
	}
	err := e.Subject.IsValid()
	if err != nil {
		return apperrors.ErrInvalidInput
	}

	switch e.Subject.Kind {
	case THEORICAL:
		//currently doesn'have a specific field
		break
	case PRACTICES:
		if e.Week == "" || e.Group == "" {
			return apperrors.ErrInvalidInput
		}

	case EXERCISES:
		if e.Group == "" {
			return apperrors.ErrInvalidInput
		}

	}
	return nil
}
