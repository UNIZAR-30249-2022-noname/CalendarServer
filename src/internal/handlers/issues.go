package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetAllIssues(c *gin.Context) {
	issues, err := hdl.Issues.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, issues)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}

//DeleteIssue is the handler for deleting a issue
//@Sumary Delete issue
//@Description Delete a issue given a id
//@Tag Issues
//@Produce text/plain
//@Param issue query string  true "id of issue"
//@Success 200 {object} string
//@Failure 400,404 {object} ErrorHttp
//@Router /deleteIssue [get]
func (hdl *HTTPHandler) DeleteIssue(c *gin.Context) {
	issue := c.Query("issue")
	err := hdl.Issues.Delete(issue)
	if err == nil {

		c.String(http.StatusOK, "Succes")
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}

//CreateIssue is the handler for creating an issue
//@Sumary Create issue
//@Description Create  a issue
//@Tag Issues
//@Produce text/plain
//@Param issue query string  true "id of issue"
//@Success 200 {object} string
//@Failure 400,404 {object} ErrorHttp
//@Router /deleteIssue [post]
func (hdl *HTTPHandler) CreateIssue(c *gin.Context) {
	body := domain.Issue{}
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorHttp{Message: "Issue en formato incorrecto"})
	}
	err = hdl.Issues.Create(body)
	if err == nil {
		c.String(http.StatusOK, "Succes")
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}

//ChangeStateIssue is the handler for changing the state of  an issue
//@Sumary Change state  issue
//@Description Create the state of a issue
//@Tag Issues
//@Produce text/plain
//@Param issue query string  true "id of issue"
//@Success 200 {object} string
//@Failure 400,404 {object} ErrorHttp
//@Router /deleteIssue [get]
func (hdl *HTTPHandler) ChangeStateIssue(c *gin.Context) {
	issue := c.Query("issue")
	newState, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorHttp{Message: "Estado en fomrato incorrecto"})
	}

	err = hdl.Issues.ChangeState(issue, newState)
	if err == nil {

		c.String(http.StatusOK, "Succes")
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}
