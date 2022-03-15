package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/gin-gonic/gin"
)

//Reserve is the handler for reserving one space
//@Sumary Reserve Space
//@Description Reserve Space a day from an initial hour to an end hour
//@Tag Users
//@Produce json
//@Param spaceId query domain.Space true "space id"
//@Param init query domain.Hour true "initial hour"
//@Param end query domain.Hour true "end hour"
//@Param date query string true "date of reserve"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /reserve/ [get]
func (hdl *HTTPHandler) Reserve(c *gin.Context) {
	id := c.Query("spaceId")
	sp := domain.Space{Id: id}
	initH, _ := strconv.Atoi(c.Query("initH"))
	initM, _ := strconv.Atoi(c.Query("initM"))
	endH, _ := strconv.Atoi(c.Query("endH"))
	endM, _ := strconv.Atoi(c.Query("endM"))
	date := c.Query("date")
	lastId, err := hdl.Spaces.Reserve(sp,domain.Hour{Hour: initH, Min: initM},domain.Hour{Hour: endH, Min: endM}, date)
	
	if err == nil {
		c.JSON(http.StatusOK, lastId)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}
}

//Reserve is the handler for reserving one space
//@Sumary Reserve Space
//@Description Reserve Space a day from an initial hour to an end hour
//@Tag Users
//@Produce json
//@Param spaces body domain.Space true "space ids"
//@Param init query domain.Hour true "initial hour"
//@Param end query domain.Hour true "end hour"
//@Param dates body []string true "dates of reserve"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /reserveBatch/ [get]
func (hdl *HTTPHandler) ReserveBatch(c *gin.Context) {
	spaces := []domain.Space{}
	c.BindJSON(&spaces)

	initH, _ := strconv.Atoi(c.Query("initH"))
	initM, _ := strconv.Atoi(c.Query("initM"))
	endH, _ := strconv.Atoi(c.Query("endH"))
	endM, _ := strconv.Atoi(c.Query("endM"))

	dates := []string{}
	c.BindJSON(&dates)
	lastId, err := hdl.Spaces.ReserveBatch(spaces,domain.Hour{Hour: initH, Min: initM},domain.Hour{Hour: endH, Min: endM}, dates)
	
	if err == nil {
		c.JSON(http.StatusOK, lastId)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}
