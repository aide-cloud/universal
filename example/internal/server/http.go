package server

import (
	"github.com/aide-cloud/universal/alog"
	"github.com/aide-cloud/universal/example/internal/conf"
	"github.com/aide-cloud/universal/example/internal/service"
	"github.com/aide-cloud/universal/p8s"
	"github.com/aide-cloud/universal/web"
	"github.com/aide-cloud/universal/web/middleware"
	"github.com/gin-gonic/gin"
)

// NewHttpServer returns a new http server.
func NewHttpServer(logger alog.Logger) *web.LierGin {
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
	r.Use(middleware.Logger(GetGlobalLog()))
	r.Use(gin.Recovery())
	r.Use(middleware.Cross())
	web.HttpPing(r, logger)
	p8s.RegisterMetricRoute(r)
	r.Use(p8s.ResponseTime())

	service.RegisterRoot(r)
}
