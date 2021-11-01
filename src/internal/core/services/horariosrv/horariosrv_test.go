package horariosrv_test

import (
	"testing"
	"time"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	mock_ports "github.com/D-D-EINA-Calendar/CalendarServer/src/mocks/mockups"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
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
	},
		{
			name: "Should return error when not found",
			args: args{terna: ternaAsked},
			want: want{result: []domain.AvailableHours{}, err: apperrors.ErrNotFound},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().GetAvailableHours(ternaAsked).Return([]domain.AvailableHours{}, apperrors.ErrNotFound)
			},
		},
		{
			name: "Should return error when [titulación] is empty",
			args: args{terna: domain.Terna{Curso: 1, Grupo: 1}},
			want: want{result: []domain.AvailableHours{}, err: apperrors.ErrInvalidInput},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().GetAvailableHours(domain.Terna{Curso: 1, Grupo: 1}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Should return error when [curso] is empty",
			args: args{terna: domain.Terna{Titulacion: "A", Grupo: 1}},
			want: want{result: []domain.AvailableHours{}, err: apperrors.ErrInvalidInput},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().GetAvailableHours(domain.Terna{Titulacion: "A", Grupo: 1}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Should return error when [curso] is empty",
			args: args{terna: domain.Terna{Titulacion: "A", Curso: 1}},
			want: want{result: []domain.AvailableHours{}, err: apperrors.ErrInvalidInput},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().GetAvailableHours(domain.Terna{Titulacion: "A", Curso: 1}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
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

			Subject:   domain.Subject{Kind: domain.THEORICAL, Name: "IC"},
			Remaining: 5,
			Max:       5,
		},
		{
			Subject:   domain.Subject{Name: "Prog 1", Kind: domain.PRACTICES},
			Remaining: 2,
			Max:       3,
		},
	}

}

/////////////////////////////
// TEST SCHEDULER ENTRIES ///
/////////////////////////////

func TestNewEntries(t *testing.T) {
	// · Mocks · //

	// · Test · //
	type args struct {
		entry domain.Entry
	}
	type want struct {
		result string
		err    error
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{{
		name: "Create entry succed",
		args: args{entry: simpleEntry()},
		want: want{result: currentDate(), err: nil},
		mocks: func(m mocks) {
			m.horarioRepository.EXPECT().SaveEntry(simpleEntry()).Return(nil)
		},
	},
		{
			name: "Should return error if repository fails",
			args: args{entry: simpleEntry()},
			want: want{result: "", err: apperrors.ErrInternal},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().SaveEntry(simpleEntry()).Return(apperrors.ErrInternal)
			},
		},
	}
	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			//Prepare
			m := mocks{
				horarioRepository: mock_ports.NewMockHorarioRepositorio(gomock.NewController(t)),
			}

			tt.mocks(m)
			service := horariosrv.New(m.horarioRepository)

			//Execute
			result, err := service.CreateNewEntry(tt.args.entry)

			//Verify operation succeded
			if tt.want.err != nil && err != nil {
				assert.Equal(t, tt.want.err.Error(), err.Error())
			}

			assert.Equal(t, tt.want.result, result)

			//Verify state changed

			//TODO use the getEntry function for verifying the entrie was created

		})

	}
}

func simpleEntry() domain.Entry {
	return domain.Entry{
		Init: domain.NewHour(1, 1),
		End:  domain.NewHour(2, 2),
		Subject: domain.Subject{
			Kind: domain.PRACTICES,
			Name: "Prog 1",
		},
		Room: domain.Room{Name: "1"},
	}
}

func currentDate() string {

	return time.Now().Format("02/01/2006")

}
