package domain

const (
	TEORIA    = 0
	PRACTICAS = 1
	PROBLEMAS = 2
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
