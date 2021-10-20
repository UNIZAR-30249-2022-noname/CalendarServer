package handlers

import (
	"net/http"
	"strconv"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
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
	if err != nil {
		//TODO hacer bien este error
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, availableHours)
}
