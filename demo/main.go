package main

import (
	"github.com/aide-cloud/universal/alog"
	"github.com/aide-cloud/universal/connect"
	"github.com/aide-cloud/universal/web/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	log := alog.NewLogger()
	r := gin.New()
	r.Use(middleware.Logger(log))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, Get(log))
	})

	if err := r.Run(":8080"); err != nil {
		return
	}
}

func Get(log alog.Logger) map[string]interface{} {
	dsn := "root:12345678@tcp(localhost:3306)/electric_app?charset=utf8&parseTime=True&loc=Local"
	db := connect.GetMysqlConnectSingle(dsn, alog.GetGormLogger(log))
	db = db.Debug()

	var map1 map[string]interface{}
	err := db.Table("electricals").Where("id > 0").Find(&map1).Error
	if err != nil {
		panic(err)
	}

	return map1
}
