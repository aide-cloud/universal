package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewCors returns a CORS middleware.
func NewCors(configs ...cors.Config) gin.HandlerFunc {
	config := cors.DefaultConfig()
	if len(configs) > 0 {
		config = configs[0]
	} else {
		config.AllowHeaders = []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"content-type-original",
			"x-requested-with",
		}
		config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
		config.AllowOriginFunc = func(origin string) bool {
			return true
		}
		config.AllowCredentials = true
	}

	return cors.New(config)
}
