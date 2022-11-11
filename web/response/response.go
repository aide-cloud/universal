package response

import "github.com/gin-gonic/gin"

func WriteJSON(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// WriteFile writes file to response
func WriteFile(c *gin.Context, data []byte, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.Data(200, "application/octet-stream", data)
}

// WriteFileHTML writes file to response
func WriteFileHTML(c *gin.Context, data []byte, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "text/html")
	c.Data(200, "text/html", data)
}

// WriteFileCSV writes file to response
func WriteFileCSV(c *gin.Context, data []byte, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "text/csv")
	c.Data(200, "text/csv", data)
}

// WriteFilePDF writes file to response
func WriteFilePDF(c *gin.Context, data []byte, fileName string) {
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/pdf")
	c.Data(200, "application/pdf", data)
}
