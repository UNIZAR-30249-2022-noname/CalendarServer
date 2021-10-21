package horariosrv_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	mock_ports "github.com/D-D-EINA-Calendar/CalendarServer/src/mocks/mockups"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

type mocks struct {
	horarioRepository *mock_ports.MockHorarioRepositorio
}

//Comprueba si dada una [Terna] devuelve las horas disponibles correctamente
func TestGetAvailableHours(t *testing.T) {
	// · Mocks · //
	AvailableHours := simpleAvailableHours()
	ternaAsked := domain.Terna{
		Titulacion: "Ing.Informática",
		Curso:      2,
		Grupo:      1,
	}

	// · Test · //
	type args struct {
		terna domain.Terna
	}
	type want struct {
		result []domain.AvailableHours
		err    error
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{{
		name: "Should return avaiable hours correctly",
		args: args{terna: ternaAsked},
		want: want{result: AvailableHours},
		mocks: func(m mocks) {
			m.horarioRepository.EXPECT().GetAvailableHours(ternaAsked).Return(AvailableHours, nil)
		},
		//TODO casos de error
	},
	}
	// · Runner · //
	for _, tt := range tests {
		//Prepare

		m := mocks{
			horarioRepository: mock_ports.NewMockHorarioRepositorio(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := horariosrv.New(m.horarioRepository)

		//Execute
		result, err := service.GetAvailableHours(tt.args.terna)

		//Verify
		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)

	}
}

func simpleAvailableHours() []domain.AvailableHours {

	return []domain.AvailableHours{
		{
			Kind:      domain.TEORIA,
			Subject:   "IC",
			Remaining: 5,
			Max:       5,
		},
		{
			Kind:      domain.PRACTICAS,
			Subject:   "Prog 1",
			Remaining: 2,
			Max:       3,
		},
	}

}
