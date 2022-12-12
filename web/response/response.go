package response

import (
	"github.com/aide-cloud/universal/aerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(c *gin.Context, data interface{}, err error) {
	if err != nil {
		switch err.(type) {
		case aerror.Error:
			c.JSON(err.(aerror.Error).HTTPStatus(), gin.H{
				"code":    err.(aerror.Error).Code(),
				"message": err.(aerror.Error).Message(),
				"data":    data,
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
				"data":    data,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

// WriteFile writes file to response
func WriteFile(c *gin.Context, data []byte, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", data)
}

// WriteFileHTML writes file to response
func WriteFileHTML(c *gin.Context, data []byte, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "text/html")
	c.Data(http.StatusOK, "text/html", data)
}

// WriteFileCSV writes file to response
func WriteFileCSV(c *gin.Context, data []byte, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "text/csv")
	c.Data(http.StatusOK, "text/csv", data)
}

// WriteFilePDF writes file to response
func WriteFilePDF(c *gin.Context, data []byte, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/pdf")
	c.Data(http.StatusOK, "application/pdf", data)
}
