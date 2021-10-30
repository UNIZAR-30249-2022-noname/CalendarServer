package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	main "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/cmd"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	mock_ports "github.com/D-D-EINA-Calendar/CalendarServer/src/mocks/mockups"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	horarioService *mock_ports.MockHorarioService
}

func TestPingRoute(t *testing.T) {
	router := main.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetAvailableHours(t *testing.T) {

	// · Mocks · //
	availableHours := simpleAvailableHours()
	errorParam := handlers.ErrorHttp{Message: "Parámetros incorrectos"}
	// · Test · //
	type args struct {
		terna handlers.TernaDto
	}

	type want struct {
		result interface{}
		code   int
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should return available hours succesfully",
			args: args{terna: handlers.TernaDto{
				Titulacion: "Ing.Informática",
				Curso:      2,
				Grupo:      1,
			}},
			want: want{result: []domain.AvailableHours{availableHours}, code: http.StatusOK},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.Terna{
					Titulacion: "Ing.Informática",
					Curso:      2,
					Grupo:      1}).Return([]domain.AvailableHours{availableHours}, nil)
			},
		},
		{
			name: "Error when [Titulacion] is empty",
			args: args{terna: handlers.TernaDto{

				Curso: 2,
				Grupo: 1,
			}},
			want: want{result: errorParam, code: http.StatusBadRequest},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.Terna{
					Curso: 2,
					Grupo: 1}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Error when [curso] is empty",
			args: args{terna: handlers.TernaDto{

				Titulacion: "Ing.Informática",
				Grupo:      1,
			}},
			want: want{result: errorParam, code: http.StatusBadRequest},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.Terna{
					Titulacion: "Ing.Informática",
					Grupo:      1}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Error when [Grupo] is empty",
			args: args{terna: handlers.TernaDto{

				Titulacion: "Ing.Informática",
				Curso:      1,
			}},
			want: want{result: errorParam, code: http.StatusBadRequest},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.Terna{
					Titulacion: "Ing.Informática",
					Curso:      1}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Error [terna] has not resources attached",
			args: args{terna: handlers.TernaDto{
				Titulacion: "Ing.Informática",
				Curso:      2,
				Grupo:      1,
			}},
			want: want{result: handlers.ErrorHttp{Message: "La terna no existe"}, code: http.StatusNotFound},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.Terna{
					Titulacion: "Ing.Informática",
					Curso:      2,
					Grupo:      1}).Return([]domain.AvailableHours{}, apperrors.ErrNotFound)
			},
		},
		//TODO more tests
	}
	// · Runner · //
	for _, tt := range tests {
		//Prepare
		m := mocks{
			horarioService: mock_ports.NewMockHorarioService(gomock.NewController(t)),
		}
		tt.mocks(m)
		setUpRouter := func() *gin.Engine {
			horarioHandler := handlers.NewHTTPHandler(m.horarioService)
			r := gin.Default()
			r.GET("/availableHours", horarioHandler.GetAvailableHours)
			return r

		}
		r := setUpRouter()
		w := httptest.NewRecorder()
		uri := "/availableHours?titulacion=" + tt.args.terna.Titulacion + "&year=" + strconv.Itoa(tt.args.terna.Curso) + "&group=" + strconv.Itoa(tt.args.terna.Grupo)
		req, _ := http.NewRequest("GET", uri, nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, tt.want.code, w.Code)

		wantedJson, _ := json.Marshal(tt.want.result)
		assert.Equal(t, bytes.NewBuffer(wantedJson), w.Body)
	}

}

func simpleAvailableHours() domain.AvailableHours {
	return domain.AvailableHours{}
}

/////////////////////////////
// TEST SCHEDULER ENTRIES ///
/////////////////////////////

func TestPostSchedulerEntry(t *testing.T) {

	// · Mocks · //

	//errorParam := handlers.ErrorHttp{Message: "Parámetros incorrectos"}

	// · Test · //
	path := "/newEntry"

	type args struct {
		newEntry handlers.EntryDTO
	}

	type want struct {
		result interface{}
		code   int
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should create a new entry succesfully",
			args: args{
				newEntry: simpleEntryKindOne()},
			want: want{result: "01/01/2021", code: http.StatusOK},
			mocks: func(m mocks) {

				m.horarioService.EXPECT().CreateNewEntry(simpleEntryKindOne().ToEntry()).Return("01/01/2021", nil)
			},
		},

		//TODO more tests
	}
	// · Runner · //
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Prepare

			m := mocks{
				horarioService: mock_ports.NewMockHorarioService(gomock.NewController(t)),
			}
			setUpRouter := func() *gin.Engine {
				horarioHandler := handlers.NewHTTPHandler(m.horarioService)
				r := gin.Default()
				r.POST(path, horarioHandler.NewEntry)
				return r

			}
			tt.mocks(m)

			r := setUpRouter()
			w := httptest.NewRecorder()
			uri := path
			body, _ := json.Marshal(simpleEntryKindOne())
			bytes.NewBuffer(body)
			req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(body))
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.want.code, w.Code)

			assert.Equal(t, tt.want.result, w.Body.String())

		})

	}

}
func simpleEntryKindOne() handlers.EntryDTO {
	return handlers.EntryDTO{
		InitHour: 1,
		InitMin:  1,
		EndHour:  1,
		EndMin:   1,
		Subject:  "a",
		Kind:     1,
		Room:     "a",
	}
}
