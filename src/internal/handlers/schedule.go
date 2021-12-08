package handlers

import (
	"fmt"
	"net/http"
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
//@Param titulacion query string true "titulacion de las horas a obtener"
//@Param curso query int true "curso de las horas a obtener"
//@Param grupo query int true "grupo de las horas a obtener"
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
