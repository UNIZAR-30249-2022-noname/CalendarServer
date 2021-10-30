package handlers

import (
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

//GetAvailableHours is the handler of get available hours endpoint
//@Sumary Get available hours
//@Description List all the hours remaining for creaiting an entrie on the schedule
//@Descriptionby type of hour (lessons, lab or problems)
//@Tag Scheduler
//@Produce json
//@Param titulacion query string true "titulacion de las horas a obtener"
//@Param curso query int true "curso de las horas a obtener"
//@Param grupo query int true "grupo de las horas a obtener"
//@Success 200 {array} domain.AvailableHours
// @Failure 400,404 {object} ErrorHttp
//@Router /availableHours/ [get]
func (hdl *HTTPHandler) GetAvailableHours(c *gin.Context) {

	titulacion := c.Query("titulacion")
	curso, _ := strconv.Atoi(c.Query("year"))
	grupo, _ := strconv.Atoi(c.Query("group"))
	terna := domain.Terna{
		Curso:      curso,
		Titulacion: titulacion,
		Grupo:      grupo,
	}
	availableHours, err := hdl.horarioService.GetAvailableHours(terna)
	if err == apperrors.ErrInvalidInput {

		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorHttp{Message: "Parámetros incorrectos"})

	} else if err == apperrors.ErrNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorHttp{Message: "La terna no existe"})

	} else if err != nil {

		c.AbortWithStatusJSON(500, ErrorHttp{Message: "unkown"})
	} else {
		c.JSON(http.StatusOK, availableHours)
	}

}

func (hdl *HTTPHandler) NewEntry(c *gin.Context) {
	//Read the body request
	body := EntryDTO{}
	c.BindJSON(&body)
	//Execute service
	lastMod, err := hdl.horarioService.CreateNewEntry(body.ToEntry())
	if err == nil {
		c.String(http.StatusOK, lastMod)

	}

}
