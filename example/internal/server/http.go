package server

import (
	"github.com/aide-cloud/universal/alog"
	"github.com/aide-cloud/universal/connect"
	"github.com/aide-cloud/universal/example/internal/conf"
	"github.com/aide-cloud/universal/example/internal/service"
	"github.com/aide-cloud/universal/executor"
	"github.com/aide-cloud/universal/p8s"
	"github.com/aide-cloud/universal/web"
	"github.com/aide-cloud/universal/web/middleware"
	"github.com/gin-gonic/gin"
)

const HTTPName = "lier-web-http"

// NewHttpServer returns a new http server.
func NewHttpServer(logger alog.Logger) executor.Service {
	gin.SetMode(conf.GetConfig().Server.Mode)
	lierGin := web.NewGin(
		web.WithEngine(gin.New()),
		web.WithGinRouters(registerRouter),
		web.WithLogger(logger),
		web.WithName(HTTPName),
		web.WithGinServer(web.NewServer(
			web.WithServerAddr(conf.GetConfig().Server.HttpAddr),
			web.WithServerMaxHeaderBytes(1<<20),
		)),
	)
	return lierGin
}

// registerRouter registers the router.
func registerRouter(r *gin.Engine, log alog.Logger) {
	r.Use(middleware.Recover(log))
	r.Use(middleware.Cross())
	web.HttpPing(r, log)
	p8s.RegisterMetricRoute(r)
	r.Use(p8s.ResponseTime())

	r.GET("/test", func(ctx *gin.Context) {
		dsn := "root:12345678@tcp(localhost:3306)/electric_app?charset=utf8&parseTime=True&loc=Local"
		db := connect.GetMysqlConnectSingle("electric_app", dsn, alog.GetGormLogger(log))
		db = db.Debug()

		var map1 map[string]interface{}
		err := db.Table("electricals").Where("id > 0").Find(&map1).Error
		if err != nil {
			panic(err)
		}

		ctx.JSON(200, map1)
	})

	service.RegisterRoot(r)
}
