package horarioRepositorio_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	consultas "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/sql"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAllBeforeTest(t *testing.T) {
	repos := horarioRepositorio.New()
	repos.RawExec(consultas.TruncEntry);		repos.RawExec(consultas.TruncHora); 
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree);	repos.RawExec(consultas.TruncAula)
}

func TestGetAvaiableHours(t *testing.T) {

	//Prepare
	assert := assert.New(t)
	hoursexpected := []domain.AvailableHours{
		{
			Subject:   domain.Subject{Kind: 1,Name: "Proyecto Software"},
			RemainingHours: 29,
			MaxHours:       30,
			RemainingMin:   0,
			MaxMin:         0,
		},
		{
			Subject:        domain.Subject{Kind: 2, Name: "Sistemas Operativos"},
			RemainingHours: 25,
			MaxHours:       25,
			RemainingMin:   0,
			MaxMin:         0,
		},
	}
	ternaAsked := domain.Terna{
		Degree: "Ing. Informatica",
		Year:   1,
		Group:  "1",
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Degree1)
	repos.RawExec(consultas.Degree2)
	repos.RawExec(consultas.Year1)
	repos.RawExec(consultas.Year2)
	repos.RawExec(consultas.Asignatura1)
	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Groupdocente1)
	repos.RawExec(consultas.Groupdocente2)
	repos.RawExec(consultas.Hora1)
	repos.RawExec(consultas.Hora2)

	//Start
	hoursgot, _ := repos.GetAvailableHours(ternaAsked)

	assert.Equal(len(hoursgot), len(hoursexpected), "Should be the same length")
	for i, h := range hoursgot {
		assert.Equal(h, hoursexpected[i], "Should be the same AvaiableHours")
	}

	//Delete
	repos.RawExec(consultas.TruncHora)
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncAsignatura)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
	repos.CloseConn()
}

func TestCreateEntry(t *testing.T) {

	//Prepare
	assert := assert.New(t)
	entryAsked := domain.Entry{

		Init:    domain.NewHour(1, 30),
		End:     domain.NewHour(2, 40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "1"},
		Week:    "",
		Group:   "",
		Weekday: 1,
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Degree1)
	repos.RawExec(consultas.Degree2)
	repos.RawExec(consultas.Year1)
	repos.RawExec(consultas.Year2)
	repos.RawExec(consultas.Asignatura1)
	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Groupdocente1)
	repos.RawExec(consultas.Groupdocente2)
	repos.RawExec(consultas.Hora1)
	repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Aula1)

	//Start (Everything OK)
	repos.CreateNewEntry(entryAsked)

	assert.Equal(repos.EntryFound(entryAsked), true, "Should be the same entries")

	//Start (Empty name)
	entryAsked = domain.Entry{
		Init:    domain.NewHour(1, 30),
		End:     domain.NewHour(2, 40),
		Subject: domain.Subject{Kind: 1, Name: ""},
		Room:    domain.Room{Name: "1"},
		Week:    "",
		Group:   "",
	}

	err := repos.CreateNewEntry(entryAsked)
	if err != nil {
		assert.Equal(apperrors.ErrSql, err, "Should be the same error")
	}

	//Start (Empty room)
	entryAsked = domain.Entry{
		Init:    domain.NewHour(1, 30),
		End:     domain.NewHour(2, 40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room:    domain.Room{Name: ""},
		Week:    "",
		Group:   "",
	}

	err = repos.CreateNewEntry(entryAsked)
	if err != nil {
		assert.Equal(apperrors.ErrSql, err, "Should be the same error")
	}

	//Delete
	repos.RawExec(consultas.TruncHora)
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncAsignatura)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

//Create entry practicas & try to create entry practicas without the week and group
func TestCreateEntryPract(t *testing.T) {

	//Prepare
	assert := assert.New(t)
	entryAsked := domain.Entry{

		Init:    domain.NewHour(2, 50),
		End:     domain.NewHour(4, 50),
		Subject: domain.Subject{Kind: 2, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "2"},
		Week:    "a",
		Group:   "mananas",
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Degree1)
	repos.RawExec(consultas.Degree2)
	repos.RawExec(consultas.Year1)
	repos.RawExec(consultas.Year2)
	repos.RawExec(consultas.Asignatura1)
	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Groupdocente1)
	repos.RawExec(consultas.Groupdocente2)
	repos.RawExec(consultas.Hora1)
	repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Hora12)
	repos.RawExec(consultas.Aula1)
	repos.RawExec(consultas.Aula2)

	//Start
	repos.CreateNewEntry(entryAsked)

	assert.Equal(repos.EntryFound(entryAsked), true, "Should be the same entries")

	//Empty group
	entryAsked = domain.Entry{

		Init:    domain.NewHour(5, 30),
		End:     domain.NewHour(6, 20),
		Subject: domain.Subject{Kind: 2, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "3"},
		Week:    "a",
		Group:   "",
	}

	err := repos.CreateNewEntry(entryAsked)
	if err != nil {
		assert.Equal(apperrors.ErrInvalidKind, err, "Should be the same error")
	}

	//Empty group
	entryAsked = domain.Entry{

		Init:    domain.NewHour(5, 30),
		End:     domain.NewHour(6, 20),
		Subject: domain.Subject{Kind: 2, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "3"},
		Week:    "",
		Group:   "mananas",
	}

	err = repos.CreateNewEntry(entryAsked)
	if err != nil {
		assert.Equal(apperrors.ErrInvalidKind, err, "Should be the same error")
	}

	//Delete
	repos.RawExec(consultas.TruncHora)
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncAsignatura)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

//Create entry problemas & try to create entry problemas without the group
func TestCreateEntryProb(t *testing.T) {
	assert := assert.New(t)
	entryAsked := domain.Entry{

		Init:    domain.NewHour(5, 30),
		End:     domain.NewHour(6, 20),
		Subject: domain.Subject{Kind: 3, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "3"},
		Week:    "",
		Group:   "niapar",
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Degree1)
	repos.RawExec(consultas.Degree2)
	repos.RawExec(consultas.Year1)
	repos.RawExec(consultas.Year2)
	repos.RawExec(consultas.Asignatura1)
	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Groupdocente1)
	repos.RawExec(consultas.Groupdocente2)
	repos.RawExec(consultas.Hora1)
	repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Hora12)
	repos.RawExec(consultas.Hora13)
	repos.RawExec(consultas.Aula1)
	repos.RawExec(consultas.Aula2)
	repos.RawExec(consultas.Aula3)

	//Start (Everything Ok)
	repos.CreateNewEntry(entryAsked)

	assert.Equal(repos.EntryFound(entryAsked), true, "Should be the same entries")

	//Empty group
	entryAsked = domain.Entry{

		Init:    domain.NewHour(5, 30),
		End:     domain.NewHour(6, 20),
		Subject: domain.Subject{Kind: 3, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "3"},
		Week:    "",
		Group:   "",
	}

	err := repos.CreateNewEntry(entryAsked)
	if err != nil {
		assert.Equal(apperrors.ErrInvalidKind, err, "Should be the same error")
	}

	//Delete
	repos.RawExec(consultas.TruncHora)
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncAsignatura)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

func TestDeleteEntry(t *testing.T) {
	//Prepare
	assert := assert.New(t)
	entryAsked := domain.Entry{
		Init:    domain.NewHour(1, 30),
		End:     domain.NewHour(2, 40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "1"},
		Week:    "",
		Group:   "",
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Degree1)
	repos.RawExec(consultas.Degree2)
	repos.RawExec(consultas.Year1)
	repos.RawExec(consultas.Year2)
	repos.RawExec(consultas.Asignatura1)
	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Groupdocente1)
	repos.RawExec(consultas.Groupdocente2)
	repos.RawExec(consultas.Hora1)
	repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Aula1)

	//Start (Everything Ok)
	repos.CreateNewEntry(entryAsked)

	assert.Equal(repos.EntryFound(entryAsked), true, "Should have been deleted")

	repos.DeleteEntry(entryAsked)

	assert.Equal(repos.EntryFound(entryAsked), false, "")

	//Empty name
	entryAsked = domain.Entry{
		Init:    domain.NewHour(1, 30),
		End:     domain.NewHour(2, 40),
		Subject: domain.Subject{Kind: 1, Name: ""},
		Room:    domain.Room{Name: "1"},
		Week:    "",
		Group:   "",
	}

	err := repos.DeleteEntry(entryAsked)
	if err != nil {
		assert.Equal(apperrors.ErrSql, err, "Should be the same error")
	}

	//Empty room
	entryAsked = domain.Entry{
		Init:    domain.NewHour(1, 30),
		End:     domain.NewHour(2, 40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room:    domain.Room{Name: ""},
		Week:    "",
		Group:   "",
	}

	err = repos.DeleteEntry(entryAsked)
	if err != nil {
		assert.Equal(apperrors.ErrSql, err, "Should be the same error")
	}

	//Delete
	repos.RawExec(consultas.TruncHora)
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncAsignatura)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

//This test creates 3 entries, 2 are in the ternaAsked and 1 isn't
//it should pass if only the 2 entries in ternaAsked are deleted
func TestDeleteAllEntries(t *testing.T) {
	//Prepare

	entryAsked1 := domain.Entry{
		Init:    domain.NewHour(1, 30),
		End:     domain.NewHour(2, 40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "1"},
		Week:    "",
		Group:   "",
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Degree1)
	repos.RawExec(consultas.Degree2)
	repos.RawExec(consultas.Year1)
	repos.RawExec(consultas.Year2)
	repos.RawExec(consultas.Asignatura1)
	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Asignatura3)
	repos.RawExec(consultas.Groupdocente1)
	repos.RawExec(consultas.Groupdocente2)
	repos.RawExec(consultas.Groupdocente3)
	repos.RawExec(consultas.Hora1)
	repos.RawExec(consultas.Hora12)
	repos.RawExec(consultas.Hora3)
	repos.RawExec(consultas.Aula1)

	repos.CreateNewEntry(entryAsked1)
	assert.Equal(t, repos.EntryFound(entryAsked1), true)

	entryAsked2 := domain.Entry{
		Init:    domain.NewHour(5, 30),
		End:     domain.NewHour(6, 20),
		Subject: domain.Subject{Kind: 2, Name: "Proyecto Software"},
		Room:    domain.Room{Name: "1"},
		Week:    "a",
		Group:   "mananas",
	}

	repos.CreateNewEntry(entryAsked2)
	assert.Equal(t, repos.EntryFound(entryAsked2), true)

	entryAsked3 := domain.Entry{
		Init:    domain.NewHour(6, 30),
		End:     domain.NewHour(7, 20),
		Subject: domain.Subject{Kind: 1, Name: "Asignatura random de Mecanica"},
		Room:    domain.Room{Name: "1"},
		Week:    "",
		Group:   "",
	}

	repos.CreateNewEntry(entryAsked3)
	assert.Equal(t, repos.EntryFound(entryAsked3), true)

	ternaAsked := domain.Terna{
		Degree: "Ing. Informatica",
		Year:   1,
		Group:  "1",
	}

	//Start

	repos.DeleteAllEntries(ternaAsked)

	assert.Equal(t, repos.EntryFound(entryAsked1), false)
	assert.Equal(t, repos.EntryFound(entryAsked2), false)
	assert.Equal(t, repos.EntryFound(entryAsked3), true)

	//Delete
	repos.RawExec(consultas.TruncHora)
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncAsignatura)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

func TestListAllDegrees(t *testing.T) {

	//Prepare
	assert := assert.New(t)

	repos := horarioRepositorio.New()

	repos.RawExec(consultas.TruncHora)
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncAsignatura)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)

	//Only 1 Degree

	repos.RawExec(consultas.Degree1)
	degreeExpected := []domain.DegreeDescription{
		{Name: "Ing. Informatica"},
	}

	res, _ := repos.ListAllDegrees()
	assert.Equal(len(degreeExpected), len(res), "Should be the same length")
	for i := range res {
		assert.Equal(degreeExpected[i].Name, res[i].Name, "Should be the same Name")
		assert.Equal(len(degreeExpected[i].Groups), len(res[i].Groups), "Should be the same length")
	}

	//Degree with year

	repos.RawExec(consultas.Year1)
	degreeExpected = []domain.DegreeDescription{
		{
			Name: "Ing. Informatica",
			Groups: []domain.YearDescription{
				{Name: 1},
			},
		},
	}

	res, _ = repos.ListAllDegrees()
	assert.Equal(len(degreeExpected), len(res), "Should be the same length")
	for i := range res {
		assert.Equal(degreeExpected[i].Name, res[i].Name, "Should be the same Name")
		assert.Equal(len(degreeExpected[i].Groups), len(res[i].Groups), "Should be the same length")
		for j := range res[i].Groups {
			assert.Equal(degreeExpected[i].Groups[j].Name, res[i].Groups[j].Name, "Should be the same Name")
			assert.Equal(len(degreeExpected[i].Groups[j].Groups), len(res[i].Groups[j].Groups), "Should be the same length")
		}
	}

	//DegreeWithYearWith2Groups

	repos.RawExec(consultas.Groupdocente1)
	repos.RawExec(consultas.Groupdocente2)
	degreeExpected = []domain.DegreeDescription{
		{
			Name: "Ing. Informatica",
			Groups: []domain.YearDescription{
				{Name: 1, Groups: []string{"1", "2"}},
			},
		},
	}

	res, _ = repos.ListAllDegrees()
	assert.Equal(len(degreeExpected), len(res), "Should be the same length")
	for i := range res {
		assert.Equal(degreeExpected[i].Name, res[i].Name, "Should be the same Name")
		assert.Equal(len(degreeExpected[i].Groups), len(res[i].Groups), "Should be the same length")
		for j := range res[i].Groups {
			assert.Equal(degreeExpected[i].Groups[j].Name, res[i].Groups[j].Name, "Should be the same Name")
			assert.Equal(len(degreeExpected[i].Groups[j].Groups), len(res[i].Groups[j].Groups), "Should be the same length")
			for k := range res[j].Groups {
				assert.Equal(degreeExpected[i].Groups[j].Groups[k], res[i].Groups[j].Groups[k], "Should be the same Name")
			}
		}
	}

	//2Degree-> 1 of them With2Year -> 1 of them With2Groups
	repos.RawExec(consultas.Degree2)
	repos.RawExec(consultas.Year12)
	degreeExpected = []domain.DegreeDescription{
		{
			Name: "Ing. Informatica",
			Groups: []domain.YearDescription{
				{Name: 1, Groups: []string{"1", "2"}},
				{Name: 2},
			},
		},
		{
			Name: "Ing. Mecanica",
		},
	}

	res, _ = repos.ListAllDegrees()
	assert.Equal(len(degreeExpected), len(res), "Should be the same length")
	for i := range res {
		assert.Equal(degreeExpected[i].Name, res[i].Name, "Should be the same Name")
		assert.Equal(len(degreeExpected[i].Groups), len(res[i].Groups), "Should be the same length")
		for j := range res[i].Groups {
			assert.Equal(degreeExpected[i].Groups[j].Name, res[i].Groups[j].Name, "Should be the same Name")
			assert.Equal(len(degreeExpected[i].Groups[j].Groups), len(res[i].Groups[j].Groups), "Should be the same length")
			for k := range res[j].Groups {
				assert.Equal(degreeExpected[i].Groups[j].Groups[k], res[i].Groups[j].Groups[k], "Should be the same Name")
			}
		}
	}

	//Delete
	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree)
	repos.CloseConn()
}

//Creates 3 entries, 2 of them are in terna, 1 isn't
func TestGetEntries(t *testing.T) {

	//Prepare
	assert := assert.New(t)
	entriesexpected := []domain.Entry{
		{
			Init:    domain.NewHour(1, 30),
			End:     domain.NewHour(2, 40),
			Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
			Room:    domain.Room{Name: "1"},
			Week:    "",
			Group:   "",
			Weekday: 1,
		},
		{
			Init:    domain.NewHour(2, 50),
			End:     domain.NewHour(4, 40),
			Subject: domain.Subject{Kind: 2, Name: "Sistemas Operativos"},
			Room:    domain.Room{Name: "2"},
			Week:    "",
			Group:   "",
			Weekday: 2,
		},
	}
	ternaAsked := domain.Terna{
		Degree: "Ing. Informatica",
		Year:   1,
		Group:  "1",
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Degree1)
	repos.RawExec(consultas.Degree2)
	repos.RawExec(consultas.Year1)
	repos.RawExec(consultas.Year2)
	repos.RawExec(consultas.Asignatura1)
	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Asignatura3)
	repos.RawExec(consultas.Groupdocente1)
	repos.RawExec(consultas.Groupdocente2)
	repos.RawExec(consultas.Hora1)
	repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Aula1)
	repos.RawExec(consultas.Aula2)
	repos.RawExec(consultas.Entry1)
	repos.RawExec(consultas.Entry2)
	repos.RawExec(consultas.Entry3)
	//Start
	entriesgot, _ := repos.GetEntries(ternaAsked)

	assert.Equal(len(entriesgot), len(entriesexpected), "Should be the same length")
	for i, h := range entriesgot {
		assert.Equal(h, entriesexpected[i], "Should be the same Entries")
	}

	//Delete
	repos.RawExec(consultas.TruncEntry);		repos.RawExec(consultas.TruncHora); 
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncGroup)
	repos.RawExec(consultas.TruncYear)
	repos.RawExec(consultas.TruncDegree);	repos.RawExec(consultas.TruncAula)
	repos.CloseConn()
}
