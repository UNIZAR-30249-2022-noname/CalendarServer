package handlers

import "github.com/gin-gonic/gin"

//Ping is the handler which response with a string
//@Sumary Returns "pong"
//@Description Response "pong" if the server is currrently available
//@Tag Monitoring
//@Produce plain
//@Success 200 "Returns "pong" "
//@Router /ping/ [get]
func (hdl *HTTPHandler) Ping(c *gin.Context) {
	payload := ""
	res, err := hdl.Monitoring.Ping()
	if !res {
		payload = err.Error()
	} else {
		payload = "ame un kebab"
	}
	c.String(200, "pong "+payload)
}
