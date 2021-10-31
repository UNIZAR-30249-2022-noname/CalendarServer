package domain

const (
	THEORICAL = 0
	PRACTICES = 1
	EXERCISES = 2
)

//modelo de horas disponibles
type AvailableHours struct {
	Kind      int
	Subject   string
	Remaining int
	Max       int
}

type Terna struct {
	Titulacion string
	Curso      int
	Grupo      int
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

type Room struct {
	Name string
}

type Entry struct {
	Init    Hour
	End     Hour
	Subject AvailableHours
	Room    Room
}
