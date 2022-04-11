package handlers

import (
	"encoding/json"
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
//@Param space query string true "space id"
//@Param hour query int true "initial hour"
//@Param date query string true "date of reserve"
//@Param person query string true "person that reserves"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /reserve/ [get]
func (hdl *HTTPHandler) Reserve(c *gin.Context) {
	id := c.Query("space")

	initString := c.Query("hour")
	initInt, _ := strconv.Atoi(initString)

	init := domain.Hour{
		Hour: initInt,
		Min:  0,
	}

	end := domain.Hour{
		Hour: initInt + 1,
		Min:  0,
	}

	date := c.Query("date")
	person := c.Query("person")
	lastId, err := hdl.Spaces.Reserve(id, init, end, date, person)

	if err == nil {
		c.JSON(http.StatusOK, lastId)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}
}

//RequestInfoSlots is the handler to get all the info of a Slot and its occupation
//@Sumary Request InfoSlots
//@Description Request the info of a space and its occupation
//@Tag Users
//@Produce json
//@Param name query string true "space name or id"
//@Param date query string true "date to request"
//@Success 200 {object} domain.AllInfoSlot
//@Failure 400,404 {object} ErrorHttp
//@Router /requestInfoSlots/ [get]
func (hdl *HTTPHandler) RequestInfoSlots(c *gin.Context) {
	name := c.Query("name")
	date := c.Query("date")

	allInfo, err := hdl.Spaces.RequestInfoSlots(domain.ReqInfoSlot{Name: name, Date: date})
	if err == nil {
		fmt.Println(name + " " + date)
		c.JSON(http.StatusOK, allInfo)
	} else {
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
//@Param person query string true "person that reserves"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /reserveBatch/ [get]
func (hdl *HTTPHandler) ReserveBatch(c *gin.Context) {
	spaces := []domain.Space{}
	c.BindJSON(&spaces)

	init := domain.Hour{}
	initJSON := []byte(c.Query("init"))
	json.Unmarshal(initJSON, &init)

	end := domain.Hour{}
	endJSON := []byte(c.Query("end"))
	json.Unmarshal(endJSON, &end)

	dates := []string{}
	c.BindJSON(&dates)
	person := c.Query("person")
	lastId, err := hdl.Spaces.ReserveBatch(spaces, init, end, dates, person)

	if err == nil {
		c.JSON(http.StatusOK, lastId)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}

//FilterBy is the handler for getting spaces based on a filter
//@Sumary Get spaces
//@Description Get spaces filtered by params
//@Tag Spaces
//@Produce json
//@Param day query string  false "day of availability"
//@Param hour query domain.Hour false "hour of availability"
//@Param floor query string false "floor where is the space"
//@Param capacity query int false " minimun capacity of the space"
//@Param building query string false "building where is the space"
//@Success 200 {object} string
//@Failure 400,404 {object} ErrorHttp
//@Router /filterSlots [get]
func (hdl *HTTPHandler) FilterBy(c *gin.Context) {
	hour := domain.Hour{}
	day := c.Query("day")
	hourJSON := []byte(c.Query("hour"))
	json.Unmarshal(hourJSON, &hour)
	floor := c.Query("floor")
	capacity, _ := strconv.Atoi(c.Query("capacity"))
	building := c.Query("building")
	params := domain.SpaceFilterParams{
		Day:      day,
		Hour:     hour,
		Floor:    floor,
		Capacity: capacity,
		Building: building,
	}
	spaces, err := hdl.Spaces.FilterBy(params)
	if err == nil {

		c.JSON(http.StatusOK, spaces)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}

//CancelReserve is the handler for canceling a reserve
//@Sumary Cancel reserve
//@Description Cancel a reserve given a id
//@Tag Reserves
//@Produce string
//@Param reserve query string  true "id of reserve"
//@Success 200 {object} string
//@Failure 400,404 {object} ErrorHttp
//@Router /cancelReserve [get]
func (hdl *HTTPHandler) CancelReserve(c *gin.Context) {
	reserve := c.Query("reserve")
	err := hdl.Spaces.CancelReserve(reserve)
	if err == nil {

		c.String(http.StatusOK, "Succes")
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}

//GetReservesOwner is the handler for getting reserves per owner
//@Sumary Get reserve
//@Description Get s reserves per owner
//@Tag Reserves
//@Produce JSON
//@Param name query string  true "iname of the owner"
//@Success 200 {object} string
//@Failure 400,404 {object} ErrorHttp
//@Router /cancelReserve [get]
func (hdl *HTTPHandler) GetReservesOwner(c *gin.Context) {
	name := c.Query("name")
	reserves, err := hdl.Spaces.GetReservesOwner(name)
	if err == nil {

		c.JSON(http.StatusOK, reserves)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}
