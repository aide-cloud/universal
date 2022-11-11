package middleware

import (
	"context"
	"github.com/aide-cloud/universal/alog"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Timeout is a middleware that times out requests after a given timeout.
// context.WithTimeout(ctx, time.Duration(millisecond)*time.Millisecond)
func Timeout(timeout time.Duration, logger alog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(ctx.Request.Context(), timeout)
		ctx.Request = ctx.Request.WithContext(c)
		ch := make(chan bool)
		go func() {
			ctx.Next()
			ch <- true
		}()

		select {
		case <-c.Done():
			cancel()
			ctx.AbortWithStatus(http.StatusGatewayTimeout)
		case <-ch:
			cancel()
		}
	}
}
