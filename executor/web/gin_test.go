package web

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestLierWebGin(t *testing.T) {
	myWebServer := NewGin()
	myWebServer.SetRouters(func(router *gin.Engine) {
		router.GET("/test", func(context *gin.Context) {
			context.String(200, "hello world, test")
		})
	}, func(router *gin.Engine) {
		router.GET("/test2", func(context *gin.Context) {
			context.String(200, "hello world, test2")
		})
	})
	err := myWebServer.Start()
	if err != nil {
		t.Error(err)
	}
}

func TestLierWebGin1(t *testing.T) {
	myWebServer := NewGin()

	err := myWebServer.Start()
	if err != nil {
		t.Error(err)
	}
}
