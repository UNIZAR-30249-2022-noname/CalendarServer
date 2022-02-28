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
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	horarioService *mock_ports.MockSchedulerService
}

func TestPingRoute(t *testing.T) {
	//setup the real router
	router := main.SetupRouter()

	w := httptest.NewRecorder()
	//doing the request
	req, _ := http.NewRequest("GET", constants.PING_URL, nil)
	router.ServeHTTP(w, req)

	//Checking results
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetAvailableHours(t *testing.T) {

	// · Mocks · //
	availableHours := handlers.NewScheduler(simpleAvailableHours())
	errorParam := handlers.ErrorHttp{Message: "Parámetros incorrectos"}
	// · Test · //
	type args struct {
		terna handlers.DegreeSetDto
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
			args: args{terna: handlers.DegreeSetDto{
				Degree: "Ing.Informática",
				Year:   2,
				Group:  "1",
			}},
			want: want{result: availableHours.AvailableHours, code: http.StatusOK},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.DegreeSet{
					Degree: "Ing.Informática",
					Year:   2,
					Group:  "1"}).Return(simpleAvailableHours(), nil)
			},
		},
		{
			name: "Error when [Degree] is empty",
			args: args{terna: handlers.DegreeSetDto{

				Year:  2,
				Group: "1",
			}},
			want: want{result: errorParam, code: http.StatusBadRequest},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.DegreeSet{
					Year:  2,
					Group: "1"}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Error when [curso] is empty",
			args: args{terna: handlers.DegreeSetDto{

				Degree: "Ing.Informática",
				Group:  "1",
			}},
			want: want{result: errorParam, code: http.StatusBadRequest},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.DegreeSet{
					Degree: "Ing.Informática",
					Group:  "1"}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Error when [Group] is empty",
			args: args{terna: handlers.DegreeSetDto{

				Degree: "Ing.Informática",
				Year:   1,
			}},
			want: want{result: errorParam, code: http.StatusBadRequest},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.DegreeSet{
					Degree: "Ing.Informática",
					Year:   1}).Return([]domain.AvailableHours{}, apperrors.ErrInvalidInput)
			},
		},
		{
			name: "Error [terna] has not resources attached",
			args: args{terna: handlers.DegreeSetDto{
				Degree: "Ing.Informática",
				Year:   2,
				Group:  "1",
			}},
			want: want{result: handlers.ErrorHttp{Message: "La terna no existe"}, code: http.StatusNotFound},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetAvailableHours(domain.DegreeSet{
					Degree: "Ing.Informática",
					Year:   2,
					Group:  "1"}).Return([]domain.AvailableHours{}, apperrors.ErrNotFound)
			},
		},
	}
	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			//Prepare
			m := mocks{
				horarioService: mock_ports.NewMockSchedulerService(gomock.NewController(t)),
			}
			tt.mocks(m)
			setUpRouter := func() *gin.Engine {
				horarioHandler := handlers.NewHTTPHandler(m.horarioService, nil)
				r := gin.Default()
				r.GET(constants.GET_AVAILABLE_HOURS_URL, horarioHandler.GetAvailableHours)
				return r

			}
			r := setUpRouter()
			w := httptest.NewRecorder()
			uri := "/availableHours?degree=" + tt.args.terna.Degree +
				"&year=" + strconv.Itoa(tt.args.terna.Year) + "&group=" + tt.args.terna.Group
			req, _ := http.NewRequest("GET", uri, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.want.code, w.Code)

			wantedJson, _ := json.Marshal(tt.want.result)
			assert.Equal(t, bytes.NewBuffer(wantedJson), w.Body)

		})

	}

}

func simpleAvailableHours() []domain.AvailableHours {
	return []domain.AvailableHours{}
}

/////////////////////////////////////
// TEST UPDATE SCHEDULER  ENTRIES ///
/////////////////////////////////////

func TestPostSchedulerEntry(t *testing.T) {

	// · Mocks · //

	//errorParam := handlers.ErrorHttp{Message: "Parámetros incorrectos"}

	// · Test · //
	path := constants.UPDATE_SCHEDULER_URL

	type args struct {
		newEntry []handlers.EntryDTO
		terna    domain.DegreeSet
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
			name: "Should update  new entries succesfully with  only theorical classes",
			args: args{
				newEntry: []handlers.EntryDTO{simpleTheoricalEntry(), simpleTheoricalEntry()},
				terna:    simpleTerna()},
			want: want{result: "01/01/2021", code: http.StatusOK},
			mocks: func(m mocks) {

				m.horarioService.EXPECT().UpdateScheduler(handlers.EntriesDTOtoDomain([]handlers.EntryDTO{simpleTheoricalEntry(), simpleTheoricalEntry()}), simpleTerna()).Return("01/01/2021", nil)
			},
		},
		{
			name: "Should update  new entries succesfully with with practices classes",
			args: args{
				newEntry: []handlers.EntryDTO{simpleTheoricalEntry(), simplePracticeEntry()},
				terna:    simpleTerna()},
			want: want{result: "01/01/2021", code: http.StatusOK},
			mocks: func(m mocks) {

				m.horarioService.EXPECT().UpdateScheduler(handlers.EntriesDTOtoDomain([]handlers.EntryDTO{simpleTheoricalEntry(), simplePracticeEntry()}), simpleTerna()).Return("01/01/2021", nil)
			},
		},
		{
			name: "Should update  new entries succesfully with for exercises classes",
			args: args{
				newEntry: []handlers.EntryDTO{simpleTheoricalEntry(), simpleExercisesEntry()},
				terna:    simpleTerna()},
			want: want{result: "01/01/2021", code: http.StatusOK},
			mocks: func(m mocks) {

				m.horarioService.EXPECT().UpdateScheduler(handlers.EntriesDTOtoDomain([]handlers.EntryDTO{simpleTheoricalEntry(), simpleExercisesEntry()}), simpleTerna()).Return("01/01/2021", nil)
			},
		},
		{
			name: "Should update  succesfully with no entries",
			args: args{
				newEntry: []handlers.EntryDTO{},
				terna:    simpleTerna()},
			want: want{result: "01/01/2021", code: http.StatusOK},
			mocks: func(m mocks) {

				m.horarioService.EXPECT().UpdateScheduler([]domain.Entry{}, simpleTerna()).Return("01/01/2021", nil)
			},
		},
	}
	// · Runner · //
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Prepare

			m := mocks{
				horarioService: mock_ports.NewMockSchedulerService(gomock.NewController(t)),
			}
			setUpRouter := func() *gin.Engine {
				horarioHandler := handlers.NewHTTPHandler(m.horarioService, nil)
				r := gin.Default()
				r.POST(path, horarioHandler.PostUpdateScheduler)
				return r

			}
			tt.mocks(m)

			r := setUpRouter()
			w := httptest.NewRecorder()
			uri := path + "?degree=" + tt.args.terna.Degree +
				"&year=" + strconv.Itoa(tt.args.terna.Year) + "&group=" + tt.args.terna.Group
			body, _ := json.Marshal(tt.args.newEntry)
			bytes.NewBuffer(body)
			req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(body))
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.want.code, w.Code)

			assert.Equal(t, tt.want.result, w.Body.String())

		})

	}

}
func simpleTheoricalEntry() handlers.EntryDTO {
	return handlers.EntryDTO{
		InitHour: 1,
		InitMin:  1,
		EndHour:  1,
		EndMin:   1,
		Subject:  "a",
		Kind:     constants.THEORICAL,
		Room:     "a",
	}
}

func simplePracticeEntry() handlers.EntryDTO {
	return handlers.EntryDTO{
		InitHour: 1,
		InitMin:  1,
		EndHour:  1,
		EndMin:   1,
		Subject:  "a",
		Kind:     constants.PRACTICES,
		Room:     "a",
		Week:     "A",
		Group:    "1",
	}
}

func simpleExercisesEntry() handlers.EntryDTO {
	return handlers.EntryDTO{
		InitHour: 1,
		InitMin:  1,
		EndHour:  1,
		EndMin:   1,
		Subject:  "a",
		Kind:     constants.EXERCISES,
		Room:     "a",
		Group:    "1",
	}
}

func simpleTerna() domain.DegreeSet {
	return domain.DegreeSet{
		Group:  "1",
		Year:   1,
		Degree: "Ing Informática",
	}
}

/////////////////////////////
// TEST LIST DEGREES      ///
/////////////////////////////

func TestListDegrees(t *testing.T) {
	// · Mocks · //

	// · Test · //
	path := constants.LIST_DEGREES_URL

	type want struct {
		result interface{}
		code   int
	}
	tests := []struct {
		name  string
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Succeded",
			want: want{result: simpleListDegreeDescriptions(), code: http.StatusOK},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().ListAllDegrees().Return(simpleListDegreeDescriptions(), nil)
			},
		},

		{
			name: "Repo failure",
			want: want{result: handlers.ErrorHttp{Message: "unkown"}, code: http.StatusInternalServerError},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().ListAllDegrees().Return(nil, apperrors.ErrInternal)
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			//Prepare
			m := mocks{
				horarioService: mock_ports.NewMockSchedulerService(gomock.NewController(t)),
			}
			setUpRouter := func() *gin.Engine {
				horarioHandler := handlers.NewHTTPHandler(m.horarioService, nil)
				r := gin.Default()
				r.GET(path, horarioHandler.ListDegrees)
				return r

			}
			tt.mocks(m)

			//Execute
			r := setUpRouter()
			w := httptest.NewRecorder()
			uri := path
			req, _ := http.NewRequest("GET", uri, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.want.code, w.Code)

			//assert.Equal(t, tt.want.result, w.Body.String())

			wantedJson, _ := json.Marshal(tt.want.result)
			assert.Equal(t, bytes.NewBuffer(wantedJson), w.Body)

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

////////////////////////
// TEST GET  ENTRIES ///
///////////////////////

func TestGetEntries(t *testing.T) {
	// · Mocks · //

	// · Test · //
	path := constants.LIST_SCHEDULER_ENTRIES_URL
	type args struct {
		terna handlers.DegreeSetDto
	}
	type want struct {
		result interface{}
		code   int
	}
	tests := []struct {
		args  args
		name  string
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should return entries succesfully",
			args: args{terna: simpleTernaDTO()},
			want: want{result: simpleListEntriesDTO(), code: http.StatusOK},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetEntries(simpleTerna()).Return(handlers.EntriesDTOtoDomain(simpleListEntriesDTO()), nil)
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			//Prepare
			m := mocks{
				horarioService: mock_ports.NewMockSchedulerService(gomock.NewController(t)),
			}
			tt.mocks(m)
			setUpRouter := func() *gin.Engine {
				horarioHandler := handlers.NewHTTPHandler(m.horarioService, nil)
				r := gin.Default()
				r.GET(path, horarioHandler.GetEntries)
				return r

			}
			r := setUpRouter()
			w := httptest.NewRecorder()
			uri := path + "?degree=" + tt.args.terna.Degree +
				"&year=" + strconv.Itoa(tt.args.terna.Year) + "&group=" + tt.args.terna.Group
			req, _ := http.NewRequest("GET", uri, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.want.code, w.Code)

			wantedJson, _ := json.Marshal(tt.want.result)
			assert.Equal(t, bytes.NewBuffer(wantedJson), w.Body)

		})

	}
}

func TestGetICS(t *testing.T) {
	// · Mocks · //
	// · Test · //
	path := constants.GENERATE_ICAL_URL
	type args struct {
		terna handlers.DegreeSetDto
	}
	type want struct {
		result interface{}
		code   int
	}
	tests := []struct {
		args  args
		name  string
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should return ICS succesfully",
			args: args{terna: simpleTernaDTO()},
			want: want{result: simpleICSFormat(), code: http.StatusOK},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().GetICS(simpleTerna()).Return(simpleICSFormat(), nil)
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			//Prepare
			m := mocks{
				horarioService: mock_ports.NewMockSchedulerService(gomock.NewController(t)),
			}
			tt.mocks(m)
			setUpRouter := func() *gin.Engine {
				horarioHandler := handlers.NewHTTPHandler(m.horarioService, nil)
				r := gin.Default()
				r.GET(path, horarioHandler.GetICS)
				return r

			}
			r := setUpRouter()
			w := httptest.NewRecorder()
			uri := path + "?degree=" + tt.args.terna.Degree +
				"&year=" + strconv.Itoa(tt.args.terna.Year) + "&group=" + tt.args.terna.Group
			req, _ := http.NewRequest("GET", uri, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.want.code, w.Code)

			wantedJson, _ := json.Marshal(tt.want.result)
			assert.Equal(t, bytes.NewBuffer(wantedJson), w.Body)

		})

	}
}

//TODO fix it
/*
//The argument isn't a string anymore
func TestUpdateByCSV(t *testing.T) {
	t.Skip("This isn't a csv anymore but we will do it properly")
	// · Mocks · //
	//content, _ := ioutil.ReadFile("../../pkg/csv/Listado207_1Asig.csv") //no cabe
	// · Test · //
	path := constants.UPLOAD_DATA_DEGREES_URL
	type args struct {
		csv string
	}
	type want struct {
		result interface{}
		code   int
	}
	tests := []struct {
		args  args
		name  string
		want  want
		mocks func(m mocks)
	}{
		{
			name: "Should return ICS succesfully",
			args: args{csv: ""},
			want: want{result: true, code: http.StatusOK},
			mocks: func(m mocks) {
				m.horarioService.EXPECT().UpdateByCSV("").Return(true, nil)
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			//Prepare
			m := mocks{
				uploadDataService: mock_ports.NewMockSchedulerService(gomock.NewController(t)),
			}
			tt.mocks(m)
			setUpRouter := func() *gin.Engine {
				horarioHandler := handlers.NewHTTPHandler(nil, m.uploadDataService)
				r := gin.Default()
				r.POST(path, horarioHandler.UpdateByCSV)
				return r

			}
			r := setUpRouter()
			w := httptest.NewRecorder()
			uri := path + "?csv=" + tt.args.csv
			req, _ := http.NewRequest("POST", uri, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.want.code, w.Code)

			wantedJson, _ := json.Marshal(tt.want.result)
			assert.Equal(t, bytes.NewBuffer(wantedJson), w.Body)

		})

	}
}
*/
func simpleTernaDTO() handlers.DegreeSetDto {
	return handlers.DegreeSetDto(simpleTerna())
}

/*
func simpleTerna() domain.DegreeSet{
	return domain.DegreeSet{
		Degree: "Ing.Informática",
		Year:      2,
		Group:      "1",
	}
}*/

func simpleICSFormat() string {
	return "BEGIN:VCALENDAR\r\nVERSION:2.0\r\nPRODID:-//arran4//Golang ICS Library\r\nBEGIN:VEVENT\r\nUID:0@unizar.es\r\nSUMMARY:Proyecto Software\r\nDTSTART:20220208T110000Z\r\nDTEND:20220208T120000Z\r\nRRULE:FREQ=DAILY;INTERVAL=7\r\nEND:VEVENT\r\nBEGIN:VEVENT\r\nUID:1@unizar.es\r\nSUMMARY:Sistemas Operativos\r\nDTSTART:20220209T090000Z\r\nDTEND:20220209T110000Z\r\nRRULE:FREQ=DAILY;INTERVAL=7\r\nEND:VEVENT\r\nBEGIN:VEVENT\r\nUID:2@unizar.es\r\nSUMMARY:Proyecto Software\r\nDTSTART:20220210T140000Z\r\nDTEND:20220210T160000Z\r\nRRULE:FREQ=DAILY;INTERVAL=7\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n"
}

func simpleListEntriesDTO() []handlers.EntryDTO {

	return []handlers.EntryDTO{simpleExercisesEntry(), simpleTheoricalEntry()}
}
