package handlers

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//UpdateByCSV is the handler for updating the database via CSV
//@Sumary Post update by CSV
//@Description The request will update the database creating degrees, subjects, years, groups and hours
//@Tag Scheduler
//@Param csv formData file true "csv file"
//@Produce json
//@Success 200 {object} bool
//@Failure 400,404 {object} ErrorHttp
//@Router /updateByCSV/ [post]
func (hdl *HTTPHandler) UpdateByCSV(c *gin.Context) {
	//Thank you! https://github.com/Cyantosh0/go-csv
	type csvUploadInput struct {
		CsvFile *multipart.FileHeader `form:"csv" binding:"required"`
	}

	var input csvUploadInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else if filepath.Ext(input.CsvFile.Filename) != ".csv" && input.CsvFile.Header.Get("Content-Type") != "text/csv" {
		c.JSON(400, gin.H{"error": "upload a csv file"})
		return
	}

	f, err := input.CsvFile.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	fileBytes, _ := ioutil.ReadAll(f)
	success, err := hdl.uploadDataservice.UpdateByCSV(string(fileBytes))
	if err == nil {
		c.JSON(http.StatusOK, success)
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}
}
