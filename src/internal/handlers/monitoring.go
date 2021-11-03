package handlers

import "github.com/gin-gonic/gin"

//Ping is the handler which response with a string
//@Sumary Returns "pong"
//@Description Response "pong" if the server is currrently available
//@Tag Monitoring
//@Produce plain
//@Success 200 "Returns "pong" "
//@Router /ping/ [get]
func Ping(c *gin.Context) {
	c.String(200, "pong")
}
