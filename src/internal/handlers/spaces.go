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
//@Param spaceId query domain.Space true "space id"
//@Param init query domain.Hour true "initial hour"
//@Param end query domain.Hour true "end hour"
//@Param date query string true "date of reserve"
//@Param person query string true "person that reserves"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /reserve/ [get]
func (hdl *HTTPHandler) Reserve(c *gin.Context) {
	id := c.Query("spaceId")
	sp := domain.Space{Id: id}

	init := domain.Hour{}
	initJSON := []byte(c.Query("init"))
	json.Unmarshal(initJSON, &init)

	end := domain.Hour{}
	endJSON := []byte(c.Query("end"))
	json.Unmarshal(endJSON, &end)

	date := c.Query("date")
	person := c.Query("person")
	lastId, err := hdl.Spaces.Reserve(sp ,init , end, date, person)
	
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

	sd := domain.SlotData{
		Name: "A1",
		Capacity: 5,
		Description: "Lorem ipsum no leas mas porque esto es dummy text",
		Building: "Ada",
		Floor: "baja",
  	};
  
    is := []domain.InfoSlots{
		{
		Hour: 8,
		Occupied: false,
		},
		{
			Hour: 8,
			Occupied: true,
			Person: "Urrikote",
		},
		{
			Hour: 10,
			Occupied: false,
		},
		{
			Hour: 11,
			Occupied: false,
		},
		{
			Hour: 12,
			Occupied: true,
			Person: "Urrikyu",
		},
		{
			Hour: 13,
			Occupied: false,
		},
		{
			Hour: 14,
			Occupied: false,
		},
		{
			Hour: 15,
			Occupied: true,
			Person: "Urriuuuu",
		},
		{
			Hour: 16,
			Occupied: false,
		},
		{
			Hour: 17,
			Occupied: false,
		},
		{
			Hour: 8,
			Occupied: true,
			Person: "Urrikoncio",
		},
		{
			Hour: 19,
			Occupied: false,
		},
		{
			Hour: 20,
			Occupied: false,
		},
	}

	allInfo := domain.AllInfoSlot{
		SlotData: sd,
		InfoSlots: is,
	}

	if( name != "" && date != ""){
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
	lastId, err := hdl.Spaces.ReserveBatch(spaces,init,end,dates,person)
	
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
	credentials, err := hdl.Spaces.FilterBy(params)
	if err == nil {

		c.JSON(http.StatusOK, credentials)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}
