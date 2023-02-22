package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode = 0
	FailedCode  = 500
)

var (
	MsgMap = map[int]string{
		SuccessCode: "success",
		FailedCode:  "failed",
	}
)

type (
	Interface interface {
		Success(c *gin.Context, msg string, data interface{})
		Failed(c *gin.Context, msg string, err error)
	}

	defaultResp struct{}
)

func NewDefault() Interface {
	return &defaultResp{}
}

func (l *defaultResp) Success(c *gin.Context, msg string, data interface{}) {
	Success(c, msg, data)
	c.Abort()
}

func (l *defaultResp) Failed(c *gin.Context, msg string, err error) {
	Failed(c, msg, err)
	c.Abort()
}

func getErrorMsg(msg string) string {
	if msg == "" {
		return MsgMap[FailedCode]
	}
	return msg
}

func getSuccessMsg(msg string) string {
	if msg == "" {
		return MsgMap[SuccessCode]
	}
	return msg
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": SuccessCode,
		"msg":  getSuccessMsg(msg),
		"data": data,
	})
}

func Failed(c *gin.Context, msg string, err error) {
	if err == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"code": FailedCode,
		"msg":  getErrorMsg(msg),
		"err":  err.Error(),
	})
}
