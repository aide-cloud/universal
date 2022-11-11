package server

import (
	"github.com/aide-cloud/universal/alog"
	"github.com/aide-cloud/universal/example/internal/conf"
	"github.com/aide-cloud/universal/example/internal/service"
	"github.com/aide-cloud/universal/executor"
	"github.com/aide-cloud/universal/p8s"
	"github.com/aide-cloud/universal/web"
	"github.com/aide-cloud/universal/web/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

// NewHttpServer returns a new http server.
func NewHttpServer(logger alog.Logger) executor.Service {
	gin.SetMode(conf.GetConfig().Server.Mode)
	lierGin := web.NewGin(
		web.WithEngine(gin.New()),
		web.WithGinRouters(registerRouter),
		web.WithLogger(logger),
		web.WithName("http"),
		web.WithGinServer(web.NewServer(
			web.WithServerAddr(conf.GetConfig().Server.HttpAddr),
			web.WithServerMaxHeaderBytes(1<<20),
		)),
	)
	return lierGin
}

// registerRouter registers the router.
func registerRouter(r *gin.Engine, logger alog.Logger) {
	r.Use(middleware.Recover(logger))
	r.Use(middleware.Timeout(time.Second*1, logger))
	r.Use(middleware.Cross())
	r.GET("/", func(ctx *gin.Context) {
		logger.Info("Hello World-1")
		for i := 0; i < 5; i++ {
			logger.Info("sleep 1s")
			time.Sleep(time.Second * 1)
		}
		logger.Info("Hello World-2")
		if ctx.Err() != nil {
			logger.Info("ctx.Err() != nil")
			return
		}
		// 判断响应是否超过写入大小
		if ctx.Writer.Size() > 0 {
			logger.Info("size > 0")
			return
		}
		logger.Info("Hello World-3")
		ctx.String(200, "Hello World")
	})
	web.HttpPing(r, logger)
	p8s.RegisterMetricRoute(r)
	r.Use(p8s.ResponseTime())

	service.RegisterRoot(r)
}
