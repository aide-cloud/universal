package web

import (
	"context"
	"fmt"
	"github.com/aide-cloud/universal/alog"
	"github.com/aide-cloud/universal/executor"
	"github.com/aide-cloud/universal/web/middleware"
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
		engine             *gin.Engine
		server             *Server
		registerRouterFunc []Router
		log                alog.Logger
		name               string
	}

	LierGinOption func(*LierGin)
)

func (l *LierGin) Name() string {
	return l.name
}

var _ executor.Service = (*LierGin)(nil)

func NewGin(options ...LierGinOption) *LierGin {
	l := &LierGin{}
	for _, option := range options {
		option(l)
	}

	if l.log == nil {
		WithLogger(alog.NewLogger())(l)
	}

	if l.engine == nil {
		WithEngine(gin.New())(l)
	}

	if l.server == nil {
		l.SetServer(NewServer(WithServerAddr(fmt.Sprintf("%s:%d", LOCALHOST, PORT8080))))
	}

	if len(l.registerRouterFunc) == 0 {
		l.SetRouters(HttpPing, httpSelfIntroduction)
	}

	r := l.engine

	r.Use(middleware.Logger(l.log))

	for _, f := range l.registerRouterFunc {
		f(r)
	}

	l.server.Handler = l.engine

	return l
}

// WithEngine set server handler
func WithEngine(engine *gin.Engine) LierGinOption {
	return func(l *LierGin) {
		l.engine = engine
	}
}

func WithGinServer(c *Server) LierGinOption {
	return func(l *LierGin) {
		l.SetServer(c)
	}
}

func WithGinRouters(routerFunc ...Router) LierGinOption {
	return func(l *LierGin) {
		l.SetRouters(routerFunc...)
	}
}

// WithLogger set logger
func WithLogger(logger alog.Logger) LierGinOption {
	return func(l *LierGin) {
		l.log = logger
	}
}

// WithName set service name
func WithName(name string) LierGinOption {
	return func(l *LierGin) {
		l.name = name
	}
}

func (l *LierGin) SetServer(c *Server) {
	l.server = c
}

func (l *LierGin) SetRouters(routerFunc ...Router) {
	l.registerRouterFunc = append(l.registerRouterFunc, routerFunc...)
}

func (l *LierGin) Start() error {
	l.log.Info(fmt.Sprintf("listen addr: %s", "http://"+l.server.Addr))
	return l.server.ListenAndServe()
}

// Stop http server stop
func (l *LierGin) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := l.server.Shutdown(ctx)
	if err != nil {
		l.log.Error("http server shutdown error: ", alog.Arg{Key: "error", Value: err})
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
