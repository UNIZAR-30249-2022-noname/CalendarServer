package horarioRepositorio_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio"
	"github.com/stretchr/testify/assert"
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
	allTrue := true
	//TODO no sudes del error pero lo pongo asi para k no de errores
	hoursgot, _ := repos.GetAvailableHours(ternaAsked)
	for i, h := range hoursgot {
		if h != hoursexpected[i] {
			t.Errorf("Expected: %v, got: %v", hoursexpected, hoursgot)
			allTrue = false
		}
	}
	assert.Equal(t, true, allTrue)
	repos.CloseConn()
}
