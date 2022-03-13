package domain

import (
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
)

//AvaialableHours is a struct which represents the available hours
//per [Terna]
type AvailableHours struct {
	Subject        Subject
	RemainingHours int
	MaxHours       int
	RemainingMin   int
	MaxMin         int
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

//DegreeDescription is a struct which represent the charactersitics
//of a specific degree, it fields are its name and the grups it has
type DegreeDescription struct {
	Name   string            `json:"name"`
	Groups []YearDescription `json:"years"`
}

//YearDescription is a struct whicjh has the info about a certain year in a degree.
//This type has no sense alone, it must me in a [DegreeDescription]
type YearDescription struct {
	Name   int      `json:"name"`
	Groups []string `json:"groups"`
}

//Set is a struct which represent the relation among
// degrees, year and group
type DegreeSet struct {
	Degree string
	Year   int
	Group  string
}

type Hour struct {
	Hour int
	Min  int
}

type Space struct {
	
}

func NewHour(h, m int) Hour {
	return Hour{
		Hour: h,
		Min:  m,
	}
}

func (h Hour) IsLaterThan(h2 Hour) bool {
	//if the hour is previus return false
	if h.Hour < h2.Hour {
		return false
		//if the hour is equal check the minutes
	} else if h.Hour == h2.Hour && h.Min <= h2.Min {
		return false
	}
	return true
}

//Los minutos no pueden pasar de 60
//Una hora y media devolveria 130
func HourToInt(h Hour) int {
	return h.Hour*100 + h.Min
}

func IntToHour(h int) Hour {
	return Hour{Hour: h / 100, Min: h % 100}
}

func AddHour(h1, h2 Hour) int {
	mins := h1.Min + h2.Min
	hours := h1.Hour + h2.Hour
	if mins >= 60 {
		mins -= 60
		hours += 1
	}
	return hours*100 + mins
}

func SubstractHour(h1, h2 Hour) int {
	mins := h1.Min - h2.Min
	hours := h1.Hour - h2.Hour
	if mins < 0 {
		mins += 60
		hours -= 1
	}
	return hours*100 + mins
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
	Weekday int
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
	case constants.THEORICAL:
		//currently doesn'have a specific field
		break
	case constants.PRACTICES:
		if e.Week == "" || e.Group == "" {
			return apperrors.ErrInvalidInput
		}

	case constants.EXERCISES:
		if e.Group == "" {
			return apperrors.ErrInvalidInput
		}

	}
	return nil
}

type User struct {
	Name       string `json:"name"`
	Privileges string `json:"privileges"`
}
