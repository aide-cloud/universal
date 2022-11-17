package server

import (
	"github.com/aide-cloud/universal/alog"
	"github.com/aide-cloud/universal/connect"
	"github.com/aide-cloud/universal/example/internal/conf"
	"github.com/aide-cloud/universal/example/internal/service"
	"github.com/aide-cloud/universal/executor"
	"github.com/aide-cloud/universal/web"
	"github.com/aide-cloud/universal/web/middleware"
	"github.com/aide-cloud/universal/web/routes"
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
	dsn := "root:12345678@tcp(localhost:3306)/electric_app1?charset=utf8&parseTime=True&loc=Local"
	connect.GetMysqlConnectSingle("electric_app", dsn, log)
	r.Use(middleware.Recover(log))
	r.Use(middleware.Cross())
	web.HttpPing(r, log)
	routes.RegisterMetricRoute(r)
	r.Use(middleware.ResponseTime())

	r.GET("/test", func(ctx *gin.Context) {
		db := connect.GetMysqlConnectSingle("electric_app", dsn, log)
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
