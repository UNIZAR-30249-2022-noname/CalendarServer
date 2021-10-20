package main_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	main "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/cmd"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	mock_ports "github.com/D-D-EINA-Calendar/CalendarServer/src/mocks/mockups"
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
	// · Test · //
	type args struct {
		terna handlers.TernaDto
	}

	type want struct {
		result []domain.AvaiableHours
		err    error
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
			want: want{result: []domain.AvaiableHours{availableHours}},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.Terna{
					Titulacion: "Ing.Informática",
					Curso:      2,
					Grupo:      1}).Return([]domain.AvaiableHours{availableHours}, nil)
			},
		},
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
		assert.Equal(t, 200, w.Code)
		//TODO comprobar horas disponibles devueltas
	}

}

func simpleAvailableHours() domain.AvaiableHours {
	return domain.AvaiableHours{}
}
