package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	LOCALHOST = "localhost"
	PORT8080  = 8080
)

type (
	Router func(router *gin.Engine)

	LierGin struct {
		server             *Server
		registerRouterFunc []Router
	}
)

func NewGin(routerFunc ...Router) *LierGin {
	l := &LierGin{}
	l.SetRouters(routerFunc...)
	return l
}

func (l *LierGin) SetServer(c *Server) {
	l.server = c
}

func (l *LierGin) SetRouters(routerFunc ...Router) {
	l.registerRouterFunc = append(l.registerRouterFunc, routerFunc...)
}

func (l *LierGin) initGin() {
	// gin初始化
	router := gin.New()

	if l.server == nil {
		l.SetServer(NewServer(WithServerAddr(fmt.Sprintf("%s:%d", LOCALHOST, PORT8080))))
	}

	if len(l.registerRouterFunc) > 0 {
		for _, f := range l.registerRouterFunc {
			f(router)
		}
	} else {
		HttpPing(router)
		httpSelfIntroduction(router)
	}

	l.server.Handler = router
}

func (l *LierGin) Start() error {
	l.initGin()
	return l.server.ListenAndServe()
}

// Stop http server stop
func (l *LierGin) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := l.server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}

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

// httpSelfIntroduction 自我介绍
func httpSelfIntroduction(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "hello world",
			"version": gin.Version,
			"now":     time.Now().Format("2006-01-02 15:04:05"),
			"author":  "biao.hu",
		})
	})
}