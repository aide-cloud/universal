package middleware

import (
	"github.com/aide-cloud/universal/alog"
	"github.com/gin-gonic/gin"
	"net/http"
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
		go func() {
			latency := time.Since(t)
			status := c.Writer.Status()
			clientIP := c.ClientIP()
			method := c.Request.Method
			errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
			args := []alog.Arg{{Key: "path", Value: path},
				{Key: "raw", Value: raw},
				{Key: "status", Value: status},
				{Key: "latency", Value: latency},
				{Key: "clientIP", Value: clientIP},
				{Key: "method", Value: method},
				{Key: "errorMessage", Value: errorMessage},
			}

			// 根据status打印日志
			switch {
			case status >= http.StatusInternalServerError:
				logger.Error("gin server error", args...)
			case status >= http.StatusBadRequest:
				logger.Warn("client error", args...)
			default:
				logger.Info("success", args...)
			}
		}()
	}
}
