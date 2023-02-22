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
