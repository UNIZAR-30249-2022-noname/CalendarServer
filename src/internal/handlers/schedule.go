package handlers

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	horarioService ports.HorarioService
}

func NewHTTPHandler(horarioService ports.HorarioService) *HTTPHandler {
	return &HTTPHandler{
		horarioService: horarioService,
	}
}

//GetAvailableHours is the handler for getting available hours endpoint
//@Sumary Get available hours
//@Description List all the hours remaining for creaiting an entrie on the schedule
//@Descriptionby type of hour (lessons, lab or problems)
//@Tag Scheduler
//@Produce json
//@Param titulacion query string true "titulacion de las horas a obtener"
//@Param curso query int true "curso de las horas a obtener"
//@Param grupo query string true "grupo de las horas a obtener"
//@Success 200 {array} domain.AvailableHours
// @Failure 400,404 {object} ErrorHttp
//@Router /availableHours/ [get]
func (hdl *HTTPHandler) GetAvailableHours(c *gin.Context) {

	titulacion := c.Query("degree")
	curso, _ := strconv.Atoi(c.Query("year"))
	grupo := c.Query("group")
	terna := domain.Terna{
		Year:   curso,
		Degree: titulacion,
		Group:  grupo,
	}
	availableHours, err := hdl.horarioService.GetAvailableHours(terna)

	if err == apperrors.ErrInvalidInput { //The set request wasn' correct

		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorHttp{Message: "Parámetros incorrectos"})

	} else if err == apperrors.ErrNotFound { //The set request does not exist
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorHttp{Message: "La terna no existe"})

	} else if err != nil { //Internal error

		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	} else {
		fmt.Println(availableHours)
		c.JSON(http.StatusOK, availableHours)

	}

}

//PostNewEntry is the handler for updating a schedluer
//@Sumary Post update a scheduler
//@Description The request will erase the current scheduler an create one new with
//@Description the requested entries for the scheduler. The entry will be definied by the initial hour
//@Description and the ending hour, adintional info must be indicated depending of the kind of hours
//@Description the kinds of subject hours are:
//@Description  - Theorical = 1
//@Description  - Practices = 2
//@Description  - Exercises = 3
//@Tag Scheduler
//@Param degree query string true "titulacion de las horas a obtener"
//@Param year query int true "curso de las horas a obtener"
//@Param group query int true "grupo de las horas a obtener"
//@Param entry body  []EntryDTO true "Entry to create"
//@Produce text/plain
//@Success 200 "Receive the date of the latests entry modification with format dd/mm/aaaa"
//@Router /updateScheduler/ [post]
func (hdl *HTTPHandler) PostUpdateScheduler(c *gin.Context) {

	titulacion := c.Query("degree")
	curso, _ := strconv.Atoi(c.Query("year"))
	grupo := c.Query("group")
	terna := domain.Terna{
		Year:   curso,
		Degree: titulacion,
		Group:  grupo,
	}
	//Read the body request
	body := []EntryDTO{}
	c.BindJSON(&body)
	listEntries := EntriesDTOtoDomain(body)

	//Execute service
	lastMod, err := hdl.horarioService.UpdateScheduler(listEntries, terna)
	if err == nil {
		c.String(http.StatusOK, lastMod)

	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}

//ListDegrees is the handler for getting the list of all degrees' descriptions avaiable
//@Sumary Get degrees description
//@Description List all degrees' descriptions avaiable, it do not require any parameter
//@Tag Scheduler
//@Produce json
//@Success 200 {array} handlers.ListDegreesDTO
// @Failure 500 {object} ErrorHttp
//@Router /listDegrees/ [get]
func (hdl *HTTPHandler) ListDegrees(c *gin.Context) {
	list, err := hdl.horarioService.ListAllDegrees()
	if err == nil {
		fmt.Println(list)
		c.JSON(http.StatusOK, list)
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}
}

//GetEntries is the handler for getting entries endpoint
//@Sumary Get entries
//@Description List all the entries of the  schedule
//@Tag Scheduler
//@Produce json
//@Param degree query string true "titulacion de las horas a obtener"
//@Param year query int true "curso de las horas a obtener"
//@Param group query int true "grupo de las horas a obtener"
//@Success 200 {array} domain.AvailableHours
// @Failure 400,404 {object} ErrorHttp
//@Router /availableHours/ [get]
func (hdl *HTTPHandler) GetEntries(c *gin.Context) {

	titulacion := c.Query("degree")
	curso, _ := strconv.Atoi(c.Query("year"))
	grupo := c.Query("group")
	terna := domain.Terna{
		Year:   curso,
		Degree: titulacion,
		Group:  grupo,
	}
	entries, err := hdl.horarioService.GetEntries(terna)

	if err == apperrors.ErrInvalidInput { //The set request wasn' correct

		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorHttp{Message: "Parámetros incorrectos"})

	} else if err == apperrors.ErrNotFound { //The set request does not exist
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorHttp{Message: "La terna no existe"})

	} else if err != nil { //Internal error

		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	} else {
		entriesDto := EntriesDomaintoDTO(entries)
		fmt.Println(entriesDto)
		c.JSON(http.StatusOK, entriesDto)

	}

}

//GetICS is the handler for getting ICalendar string endpoint
//@Sumary Get ICS
//@Description Get the schedule in ics format
//@Tag Scheduler
//@Produce json
//@Param degree query string true "titulacion de las horas a obtener"
//@Param year query int true "curso de las horas a obtener"
//@Param group query int true "grupo de las horas a obtener"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /getICS/ [get]
func (hdl *HTTPHandler) GetICS(c *gin.Context) {
	titulacion := c.Query("degree")
	curso, _ := strconv.Atoi(c.Query("year"))
	grupo := c.Query("group")
	terna := domain.Terna{
		Year:   curso,
		Degree: titulacion,
		Group:  grupo,
	}
	list, err := hdl.horarioService.GetICS(terna)
	if err == nil {
		fmt.Println(list)
		c.JSON(http.StatusOK, list)
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}
}

//UpdateByCSV is the handler for updating the database via CSV
//@Sumary Post update by CSV
//@Description The request will update the database creating degrees, subjects, years, groups and hours
//@Tag Scheduler
//@Param csv formData file true "csv file"
//@Produce json
//@Success 200 {object} bool
//@Failure 400,404 {object} ErrorHttp
//@Router /updateByCSV/ [post]
func (hdl *HTTPHandler) UpdateByCSV(c *gin.Context) {
	//Thank you! https://github.com/Cyantosh0/go-csv
	type csvUploadInput struct {
		CsvFile *multipart.FileHeader `form:"csv" binding:"required"`
	}

	var input csvUploadInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else if filepath.Ext(input.CsvFile.Filename) != ".csv" && input.CsvFile.Header.Get("Content-Type") != "text/csv" {
		c.JSON(400, gin.H{"error": "upload a csv file"})
		return
	}

	f, err := input.CsvFile.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	fileBytes, _ := ioutil.ReadAll(f)
	success, err := hdl.horarioService.UpdateByCSV(string(fileBytes))
	if err == nil {
		c.JSON(http.StatusOK, success)
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}
}
