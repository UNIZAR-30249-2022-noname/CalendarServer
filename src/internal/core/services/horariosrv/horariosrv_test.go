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

//Checks all the cases for the function GetAvailableHours of the service [horariosrv]
func TestGetAvailableHours(t *testing.T) {
	// · Mocks · //
	AvailableHours := simpleAvailableHours()
	ternaAsked := domain.Terna{
		Titulacion: "Ing.Informática",
		Curso:      2,
		Grupo:      "1",
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
			args: args{terna: domain.Terna{Curso: 1, Grupo: "1"}},
			want: want{result: []domain.AvailableHours{}, err: apperrors.ErrInvalidInput},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().GetAvailableHours(domain.Terna{Curso: 1, Grupo: "1"}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Should return error when [curso] is empty",
			args: args{terna: domain.Terna{Titulacion: "A", Grupo: "1"}},
			want: want{result: []domain.AvailableHours{}, err: apperrors.ErrInvalidInput},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().GetAvailableHours(domain.Terna{Titulacion: "A", Grupo: "1"}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
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
		t.Run(tt.name, func(t *testing.T) {
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

		})

	}
}

//Returns a slice for having different cases in the tests
func simpleAvailableHours() []domain.AvailableHours {

	return []domain.AvailableHours{
		{

			Subject:        domain.Subject{Kind: domain.THEORICAL, Name: "IC"},
			RemainingHours: 5,
			MaxHours:       5,
			RemainingMin:   0,
			MaxMin:         0,
		},
		{
			Subject:        domain.Subject{Name: "Prog 1", Kind: domain.PRACTICES},
			RemainingHours: 2,
			MaxHours:       3,
			RemainingMin:   0,
			MaxMin:         0,
		},
	}

}

/////////////////////////////////////
// TEST UPDATE SCHEDULER ENTRIES ///
///////////////////////////////////

func TestUpdateEntries(t *testing.T) {
	// · Mocks · //

	// · Test · //
	type args struct {
		entries []domain.Entry
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
		args: args{entries: simpleEntries()},
		want: want{result: currentDate(), err: nil},
		mocks: func(m mocks) {
			m.horarioRepository.EXPECT().CreateNewEntry(simpleEntries()[0]).Return(nil)
			m.horarioRepository.EXPECT().CreateNewEntry(simpleEntries()[1]).Return(nil)
			m.horarioRepository.EXPECT().DeleteAllEntries(domain.Terna{}).Return(nil) //TODO
		},
	},
		{
			name: "Should return error if repository fails",
			args: args{entries: simpleEntries()},
			want: want{result: "", err: apperrors.ErrSql},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().CreateNewEntry(simpleEntries()[0]).Return(apperrors.ErrInternal)
				m.horarioRepository.EXPECT().DeleteAllEntries(domain.Terna{}).Return(nil) //TODO
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
			result, err := service.UpdateScheduler(tt.args.entries)

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

func simpleEntries() []domain.Entry {
	return []domain.Entry{
		{
			Init: domain.NewHour(1, 1),
			End:  domain.NewHour(2, 2),
			Subject: domain.Subject{
				Kind: domain.THEORICAL,
				Name: "Prog 1",
			},
			Room:    domain.Room{Name: "1"},
			Weekday: domain.MOONDAY,
		},
		{
			Init: domain.NewHour(5, 0),
			End:  domain.NewHour(9, 0),
			Subject: domain.Subject{
				Kind: domain.THEORICAL,
				Name: "Prog 2",
			},
			Room:    domain.Room{Name: "2"},
			Weekday: domain.THUERSDAY,
		},
	}
}

func currentDate() string {

	return time.Now().Format("02/01/2006")

}

/////////////////////////////
// TEST LIST DEGREES      ///
/////////////////////////////

func TestListSubject(t *testing.T) {
	// · Mocks · //

	// · Test · //

	type want struct {
		result []domain.DegreeDescription
		err    error
	}
	tests := []struct {
		name  string
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Succeded",
			want: want{result: simpleListDegreeDescriptions(), err: nil},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().ListAllDegrees().Return(simpleListDegreeDescriptions(), nil)
			},
		},

		{
			name: "Repo failure",
			want: want{result: nil, err: apperrors.ErrInternal},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().ListAllDegrees().Return(nil, apperrors.ErrInternal)
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
			result, err := service.ListAllDegrees()

			//Verify operation succeded
			if tt.want.err != nil && err != nil {
				assert.Equal(t, tt.want.err.Error(), err.Error())
			}

			assert.Equal(t, tt.want.result, result)

		})

	}
}

func simpleListDegreeDescriptions() []domain.DegreeDescription {
	return []domain.DegreeDescription{
		{
			Name: "A",
			Groups: []domain.YearDescription{
				{Name: 1, Groups: []string{"a", "b"}},
				{Name: 2, Groups: []string{"a", "b"}},
			},
		},
		{
			Name: "B",
			Groups: []domain.YearDescription{
				{Name: 1, Groups: []string{"a"}},
				{Name: 2, Groups: []string{"a", "b", "c"}},
			},
		},
	}
}
