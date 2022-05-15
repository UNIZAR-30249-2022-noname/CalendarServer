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
//@Param slot query string true "space id"
//@Param scheduled body []domain.Hour true "initial hour"
//@Param day query string true "date of reserve"
//@Param owner query string true "person that reserves"
//@Param event query string true "event in the reserve"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /reserve/ [get]
func (hdl *HTTPHandler) Reserve(c *gin.Context) {
	body := domain.Reserve{}
	c.BindJSON(&body)

	lastId, err := hdl.Spaces.Reserve(body.Space, body.Scheduled[0], body.Scheduled[1], body.Day, body.Owner, body.Event)

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
//@Param id query string true "space id"
//@Param date query string true "date to request"
//@Success 200 {object} domain.AllInfoSlot
//@Failure 400,404 {object} ErrorHttp
//@Router /requestInfoSlots/ [get]
func (hdl *HTTPHandler) RequestInfoSlots(c *gin.Context) {
	id := c.Query("id")
	date := c.Query("date")

	fmt.Println("id: " + id + " date: " + date)
	allInfo, err := hdl.Spaces.RequestInfoSlots(domain.ReqInfoSlot{Id: id, Date: date})
	if err == nil {
		fmt.Println(id + " " + date)
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
//@Param spaces body []string true "space ids"
//@Param init query domain.Hour true "initial hour"
//@Param end query domain.Hour true "end hour"
//@Param dates body []string true "dates of reserve"
//@Param person query string true "person that reserves"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /reserveBatch/ [get]
func (hdl *HTTPHandler) ReserveBatch(c *gin.Context) {
	spaces := []string{}
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
//@Produce text/plain
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
//@Produce json
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
