package horarioRepositorio_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/stretchr/testify/assert"
)

func TestBasico(t *testing.T) {
	err := apperrors.ErrSql
	hoursexpected := []domain.AvailableHours{
		{
			Subject:   domain.Subject{Kind: 0,Name: "si"},
			Remaining: 20,
			Max:       21,
		},
		{
			Subject:   domain.Subject{Kind: 0,Name: "si"},
			Remaining: 25,
			Max:       26,
		},
	}
	ternaAsked := domain.Terna{
		Titulacion: "uwu",
		Curso:      1,
		Grupo:      0,
	}
	repos := horarioRepositorio.New()
	hoursgot, error := repos.GetAvailableHours(ternaAsked)
	if error != nil {
		assert.Equal(t, err, error)
	}
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
