package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Timeout is a middleware that times out requests after a given timeout.
// context.WithTimeout(ctx, time.Duration(millisecond)*time.Millisecond)
func Timeout(millisecond int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(ctx, time.Duration(millisecond)*time.Millisecond)
		defer cancel()
		ctx.Request = ctx.Request.WithContext(c)

		go func() {
			for {
				select {
				case <-ctx.Done():
					ctx.AbortWithStatus(http.StatusRequestTimeout)
					return
				}
			}
		}()

		ctx.Next()
	}
}
