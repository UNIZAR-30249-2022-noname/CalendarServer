package uploaddata_test

import (
	"io/ioutil"
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/services/horariosrv"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestUpdateByCSV(t *testing.T) {
	// · Mocks · //

	// · Test · //
	type args struct {
		terna domain.Terna
	}
	type want struct {
		result bool
		err    error
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should Update correctly",
			args: args{},
			want: want{result: true},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto").Return(true, nil)
				m.horarioRepository.EXPECT().CreateNewYear(1, 558).Return(true, nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('55811','1','5581'), ('55812','2','5581')").Return(nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES ('25802','25802','Informática','558')").Return(nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `idasignatura`, `idgrupo`, `grupo`, `semana`) VALUES (NULL,'3500','3500','1','25802','55811','',''), (NULL,'1000','1000','2','25802','55811','1',''), (NULL,'1000','1000','2','25802','55811','2',''), (NULL,'1500','1500','3','25802','55811','1','a'), (NULL,'1500','1500','3','25802','55811','2','a'), (NULL,'1500','1500','3','25802','55811','3','a'), (NULL,'3500','3500','1','25802','55812','',''), (NULL,'1000','1000','2','25802','55812','1',''), (NULL,'1000','1000','2','25802','55812','2',''), (NULL,'1500','1500','3','25802','55812','1','a'), (NULL,'1500','1500','3','25802','55812','2','a'), (NULL,'1500','1500','3','25802','55812','3','a')").Return(nil)
			},
		},
		{
			name: "Subject creation fails",
			args: args{},
			want: want{result: false, err: apperrors.ErrSql},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto").Return(true, nil)
				m.horarioRepository.EXPECT().CreateNewYear(1, 558).Return(true, nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('55811','1','5581'), ('55812','2','5581')").Return(nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES ('25802','25802','Informática','558')").Return(apperrors.ErrSql)
			},
		},
		{
			name: "Hour creation fails",
			args: args{},
			want: want{result: false, err: apperrors.ErrSql},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto").Return(true, nil)
				m.horarioRepository.EXPECT().CreateNewYear(1, 558).Return(true, nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('55811','1','5581'), ('55812','2','5581')").Return(nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES ('25802','25802','Informática','558')").Return(nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `idasignatura`, `idgrupo`, `grupo`, `semana`) VALUES (NULL,'3500','3500','1','25802','55811','',''), (NULL,'1000','1000','2','25802','55811','1',''), (NULL,'1000','1000','2','25802','55811','2',''), (NULL,'1500','1500','3','25802','55811','1','a'), (NULL,'1500','1500','3','25802','55811','2','a'), (NULL,'1500','1500','3','25802','55811','3','a'), (NULL,'3500','3500','1','25802','55812','',''), (NULL,'1000','1000','2','25802','55812','1',''), (NULL,'1000','1000','2','25802','55812','2',''), (NULL,'1500','1500','3','25802','55812','1','a'), (NULL,'1500','1500','3','25802','55812','2','a'), (NULL,'1500','1500','3','25802','55812','3','a')").Return(apperrors.ErrSql)
			},
		},
		{
			name: "Group creation fails",
			args: args{},
			want: want{result: false, err: apperrors.ErrSql},
			mocks: func(m mocks) {
				m.horarioRepository.EXPECT().CreateNewDegree(558, "Graduado en Ingeniería en Diseño Industrial y Desarrollo de Producto").Return(true, nil)
				m.horarioRepository.EXPECT().CreateNewYear(1, 558).Return(true, nil)
				m.horarioRepository.EXPECT().RawExec("INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('55811','1','5581'), ('55812','2','5581')").Return(apperrors.ErrSql)
			},
		},
	}

	// · Runner · //
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Prepare

			m := mocks{
				horarioRepository: mock_ports.NewMockHorarioRepositorio(gomock.NewController(t)),
			}

			tt.mocks(m)
			service := horariosrv.New(m.horarioRepository)
			content, err := ioutil.ReadFile("../../../../pkg/csv/Listado207_1Asig.csv")
			contentString := string(content)
			//Execute
			result, err := service.UpdateByCSV(contentString)

			//Verify
			if tt.want.err != nil && err != nil {
				assert.Equal(t, tt.want.err.Error(), err.Error())
			}

			if i != 0 {
				assert.Equal(t, tt.want.result, result)
			} else {
				assert.NotEqual(t, "", result)
			}

		})

	}

}
