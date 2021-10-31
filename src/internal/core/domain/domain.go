package domain

const (
	TEORIA    = 0
	PRACTICAS = 1
	PROBLEMAS = 2
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
