package middleware

import (
	"github.com/aide-cloud/universal/alog"
	"github.com/gin-gonic/gin"
)

// Recover is a middleware that recovers from any panics and writes a 500 if there was one.
func Recover(logger alog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic", alog.Arg{Key: "err", Value: err})
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
