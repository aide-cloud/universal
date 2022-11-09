package middleware

import (
	"github.com/aide-cloud/universal/alog"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger(logger alog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		// after request
		latency := time.Since(t)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		logger.Info("gin",
			alog.Arg{Key: "path", Value: path},
			alog.Arg{Key: "raw", Value: raw},
			alog.Arg{Key: "status", Value: status},
			alog.Arg{Key: "latency", Value: latency},
			alog.Arg{Key: "clientIP", Value: clientIP},
			alog.Arg{Key: "method", Value: method},
			alog.Arg{Key: "errorMessage", Value: errorMessage},
		)
	}
}
