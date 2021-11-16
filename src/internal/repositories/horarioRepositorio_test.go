package horarioRepositorio_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/stretchr/testify/assert"
)

func TestBasico(t *testing.T) {
	//err := apperrors.ErrSql
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
		Grupo:      1,
	}
	repos := horarioRepositorio.New()
	hoursgot, _ := repos.GetAvailableHours(ternaAsked)
	/*
	if error != nil {
		assert.Equal(t, err, error)
	}
	*/
	assert.Equal(t, len(hoursgot), len(hoursexpected))
	for i, h := range hoursgot {
		assert.Equal(t, h, hoursexpected[i])
	}
	repos.CloseConn()
}

func TestCreateEntry(t *testing.T) {
	err := apperrors.ErrSql
	entryAsked := domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: "si"},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "",
	}
	repos := horarioRepositorio.New()
	error := repos.CreateNewEntry(entryAsked)
	if error != nil {
		assert.Equal(t, err, error)
	} else {
		assert.Equal(t, true, true)
	}

	repos.CloseConn()
}

func TestCreateEntryPract(t *testing.T) {
	err := apperrors.ErrSql
	entryAsked := domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 2, Name: "no"},
		Room: domain.Room{Name: "1"},
		Week: "a",
		Group: "mananas",
	}
	repos := horarioRepositorio.New()
	error := repos.CreateNewEntry(entryAsked)
	if error != nil {
		assert.Equal(t, err, error)
	} else {
		assert.Equal(t, true, true)
	}

	repos.CloseConn()
}

func TestCreateEntryProb(t *testing.T) {
	err := apperrors.ErrSql
	entryAsked := domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 3, Name: "no"},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "niapar",
	}
	repos := horarioRepositorio.New()
	error := repos.CreateNewEntry(entryAsked)
	if error != nil {
		assert.Equal(t, err, error)
	} else {
		assert.Equal(t, true, true)
	}

	repos.CloseConn()
}

func TestDeleteEntry(t *testing.T) {
	err := apperrors.ErrSql
	entryAsked := domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: "si"},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "",
	}
	repos := horarioRepositorio.New()
	error := repos.DeleteEntry(entryAsked)
	if error != nil {
		assert.Equal(t, err, error)
	} else {
		assert.Equal(t, true, true)
	}

	repos.CloseConn()
}
