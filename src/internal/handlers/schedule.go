package handlers

import (
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

}
