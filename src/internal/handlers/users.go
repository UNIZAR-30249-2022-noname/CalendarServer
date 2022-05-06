package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Login is the handler for getting credentials
//@Sumary Get credentials
//@Description Get credentials for doing task which requires certain privileges
//@Tag Users
//@Produce json
//@Param username query string true "name of the user"
//@Success 200 {object} string
// @Failure 400,404 {object} ErrorHttp
//@Router /login/ [get]
func (hdl *HTTPHandler) Login(c *gin.Context) {
	username := c.Query("username")
	credentials, err := hdl.Users.GetCredentials(username)
	if err == nil {

		c.JSON(http.StatusOK, credentials)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}

}
