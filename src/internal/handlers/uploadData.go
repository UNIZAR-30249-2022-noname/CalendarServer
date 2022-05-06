package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UpdateByCSV is the handler for updating the database via CSV
//@Sumary Post update by CSV
//@Description The request will update the database creating degrees, subjects, years, groups and hours
//@Tag Scheduler
//@Param csv body string true "csv"
//@Produce json
//@Success 200 {object} bool
//@Failure 400,404 {object} ErrorHttp
//@Router /updateByCSV/ [post]
func (hdl *HTTPHandler) UpdateByCSV(c *gin.Context) {
	//Thank you! https://github.com/Cyantosh0/go-csv
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	csv := buf.String()
	fmt.Println(csv)
	success, err := hdl.UploadData.UpdateByCSV(string(csv))
	if err == nil {
		c.JSON(http.StatusOK, success)
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}
}
