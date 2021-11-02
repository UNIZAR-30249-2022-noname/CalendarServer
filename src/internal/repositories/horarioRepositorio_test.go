package horarioRepositorio_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
)

func TestBasico(t *testing.T) {
	hoursexpected := []domain.AvailableHours{
		{
			Kind:      0,
			Subject:   "si",
			Remaining: 20,
			Max:       21,
		},
		{
			Kind:      0,
			Subject:   "si",
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
	hoursgot := repos.GetAvailableHours(ternaAsked)
	if hoursgot != hoursexpected {
		t.Errorf("Expected: %v, got: %v", hoursexpected, hoursgot)
	}
}