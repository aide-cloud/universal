package main

import (
	"fmt"
	"github.com/aide-cloud/universal/executor"
	"github.com/gin-gonic/gin"
	"time"
)

type GinHttpServer struct {
}

func (m *GinHttpServer) Start() error {
	fmt.Println("GinHttpServer start")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go func() {
		err := r.Run(":8080")
		if err != nil {
			fmt.Println("GinHttpServer stop")
		}
	}()
	return nil
}

func (m *GinHttpServer) Stop() {
	fmt.Println("GinHttpServer stop")
}

func NewGinHttpServer() *GinHttpServer {
	return &GinHttpServer{}
}

func main() {
	executor.ExecMulSerProgram(executor.NewLierCmd(
		executor.WithServiceName("HelloWorld"),
		executor.WithProperty(map[string]string{
			"version": "1.0.0",
			"name   ": "cmd",
			"time   ": time.Now().Format("2006-01-02 15:04:05"),
			"author ": "aide-cloud",
		}),
		executor.WithServices(NewGinHttpServer()),
	))
}
