package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HttpPing router ping
func HttpPing(router *gin.Engine) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusOK)
	})
}

// HttpCheckHealth 检查服务健康
func HttpCheckHealth(router *gin.Engine) {
	router.GET("/check", func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusOK)
	})
}
