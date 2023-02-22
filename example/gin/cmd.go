package main

import (
	"github.com/aide-cloud/universal/executor"
	lGin "github.com/aide-cloud/universal/gin"
	"github.com/gin-gonic/gin"
)

func RegisterGinRouters(r gin.IRouter) {
	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get router",
		})
	})
	r.POST("/add", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "add router",
		})
	})
}

func main() {
	l := lGin.NewGin(
		gin.Default(),
		lGin.WithRegisterRouters(
			lGin.RegisterPing,
			RegisterGinRouters,
			func(r gin.IRouter) {
				r.GET("/hello", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"message": "hello get",
					})
				})
				r.POST("/hello", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"message": "hello ",
					})
				})
			},
		),
	)
	executor.ExecMulSerProgram(executor.NewLierCmd(
		executor.WithServiceName("master"),
		executor.WithProperty(map[string]string{
			"version": "1.0.0",
		}),
		executor.WithServices(l),
	))
}
