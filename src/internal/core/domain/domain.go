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
