package horarioRepositorio_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio"
	consultas "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/sql"
	"github.com/stretchr/testify/assert"
)


func TestBasico(t *testing.T) {
	
	//Prepare
	//err := apperrors.ErrSql
	assert := assert.New(t)
	hoursexpected := []domain.AvailableHours{
		{
			Subject:   domain.Subject{Kind: 1,Name: "Proyecto Software"},
			Remaining: 30,
			Max:       30,
		},
		{
			Subject:   domain.Subject{Kind: 2, Name: "Sistemas Operativos"},
			Remaining: 25,
			Max:       25,
		},
	}
	ternaAsked := domain.Terna{
		Titulacion: "Ing. Informatica",
		Curso:      1,
		Grupo:      "1",
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)

	//Start
	hoursgot, _ := repos.GetAvailableHours(ternaAsked)
	//if error != nil { assert.Equal(t, err, error)}
	assert.Equal(len(hoursgot), len(hoursexpected), "Should be the same length")
	for i, h := range hoursgot {
		assert.Equal(h, hoursexpected[i], "Should be the same AvaiableHours")
	}

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion)
	repos.CloseConn()
}


func TestCreateEntry(t *testing.T) {

	//Prepare
	//err := apperrors.ErrSql
	entryAsked := domain.Entry{

		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "",

	}
	
	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Aula1);

	//Start
	repos.CreateNewEntry(entryAsked)
	
	assert.Equal(t, repos.EntryFound(entryAsked), true)

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion);	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

func TestCreateEntryPract(t *testing.T) {

	//Prepare
	entryAsked := domain.Entry{

		Init: domain.NewHour(2,50),
		End: domain.NewHour(4,50),
		Subject: domain.Subject{Kind: 2, Name: "Proyecto Software"},
		Room: domain.Room{Name: "2"},
		Week: "a",
		Group: "mananas",

	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Hora12)
	repos.RawExec(consultas.Aula1);			repos.RawExec(consultas.Aula2)

	//Start
	repos.CreateNewEntry(entryAsked)
	
	assert.Equal(t, repos.EntryFound(entryAsked), true)

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion);	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

func TestCreateEntryProb(t *testing.T) {

	entryAsked := domain.Entry{

		Init: domain.NewHour(5,30),
		End: domain.NewHour(6,20),
		Subject: domain.Subject{Kind: 3, Name: "Proyecto Software"},
		Room: domain.Room{Name: "3"},
		Week: "",
		Group: "niapar",

	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Hora12);		repos.RawExec(consultas.Hora13);
	repos.RawExec(consultas.Aula1);			repos.RawExec(consultas.Aula2)
	repos.RawExec(consultas.Aula3)

	//Start
	repos.CreateNewEntry(entryAsked)
	
	assert.Equal(t, repos.EntryFound(entryAsked), true)

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion);	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

func TestDeleteEntry(t *testing.T) {
	//Prepare
	//err := apperrors.ErrSql
	entryAsked := domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "",
	}
	
	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Aula1);

	//Start
	repos.CreateNewEntry(entryAsked)
	
	assert.Equal(t, repos.EntryFound(entryAsked), true)

	repos.DeleteEntry(entryAsked)

	assert.Equal(t, repos.EntryFound(entryAsked), false)

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion);	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}
