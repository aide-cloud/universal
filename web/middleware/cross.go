package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Cross is a middleware to handle cross domain request.
func Cross() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.GetHeader("Origin")
		if origin == "null" || origin == "" {
			origin = "*"
		}
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", strings.Join([]string{"GET", "POST", "PUT", "DELETE", "PATCH"}, ","))
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-Requested-With, Accept, X-Token")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Max-Age", "3600")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
