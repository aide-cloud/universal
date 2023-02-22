package middleware

import (
	"github.com/gin-gonic/gin"
)

// Recover is a middleware that recovers from any panics and writes a 500 if there was one.
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
