package domain

const (
	THEORICAL = 0
	PRACTICES = 1
	EXERCISES = 2
)

//AvaialableHours is a struct which represents the available hours
//per [Terna]
type AvailableHours struct {
	Kind      int
	Subject   string
	Remaining int
	Max       int
}

//Terna is a struct which represent the relation among
// bachelors, year and group
type Terna struct {
	Titulacion string
	Curso      int
	Grupo      int
}

type hour struct {
	hour int
	min  int
}

func NewHour(h, m int) hour {
	return hour{
		hour: h,
		min:  m,
	}
}

type Room struct {
	Name string
}

type Entry struct {
	Init    hour
	End     hour
	Subject AvailableHours
	Room    Room
}
