package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type (
	// CrossOption is the option for Cross middleware.
	CrossOption func(*crossOption)
	crossOption struct {
		allowOrigin      string
		allowMethods     []string
		allowHeaders     []string
		allowCredentials bool
		allowMaxAge      int
	}
)

// Cross is a middleware to handle cross domain request.
func Cross(options ...CrossOption) gin.HandlerFunc {
	opt := &crossOption{
		allowOrigin:      "*",
		allowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		allowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With", "X-Auth-Token"},
		allowCredentials: true,
		allowMaxAge:      86400,
	}
	for _, option := range options {
		if option != nil {
			option(opt)
		}
	}
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.GetHeader("Origin")
		if origin == "null" || origin == "" {
			origin = "*"
		}
		c.Header("Access-Control-Allow-Origin", opt.allowOrigin)
		c.Header("Access-Control-Allow-Methods", strings.Join(opt.allowMethods, ","))
		c.Header("Access-Control-Allow-Headers", strings.Join(opt.allowHeaders, ","))
		c.Header("Access-Control-Allow-Credentials", fmt.Sprintf("%v", opt.allowCredentials))
		if method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Max-Age", strconv.Itoa(opt.allowMaxAge))
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// CrossAllowOrigin is the option for Cross middleware to set allow origin.
func CrossAllowOrigin(origin string) CrossOption {
	return func(opt *crossOption) {
		opt.allowOrigin = origin
	}
}

// CrossAllowMethods is the option for Cross middleware to set allow methods.
func CrossAllowMethods(methods []string) CrossOption {
	return func(opt *crossOption) {
		opt.allowMethods = methods
	}
}

// CrossAllowHeaders is the option for Cross middleware to set allow headers.
func CrossAllowHeaders(headers []string) CrossOption {
	return func(opt *crossOption) {
		opt.allowHeaders = headers
	}
}

// CrossAllowCredentials is the option for Cross middleware to set allow credentials.
func CrossAllowCredentials(credentials bool) CrossOption {
	return func(opt *crossOption) {
		opt.allowCredentials = credentials
	}
}

// CrossAllowMaxAge is the option for Cross middleware to set allow max age.
func CrossAllowMaxAge(maxAge int) CrossOption {
	return func(opt *crossOption) {
		opt.allowMaxAge = maxAge
	}
}
