package gin

import (
	"context"
	"fmt"
	"github.com/aide-cloud/universal/executor"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	LOCALHOST = "localhost"
	PORT8080  = 8080
)

type (
	Router func(router gin.IRouter)

	LierGin struct {
		engine          *gin.Engine
		server          *http.Server
		registerRouters []Router
	}

	LierGinOption func(*LierGin)
)

var _ executor.Service = (*LierGin)(nil)

func NewGin(r *gin.Engine, options ...LierGinOption) *LierGin {
	l := &LierGin{
		engine: r,
	}
	for _, option := range options {
		option(l)
	}

	if l.server != nil {
		l.server.Handler = r.Handler()
	} else {
		l.server = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", LOCALHOST, PORT8080),
			Handler: r.Handler(),
		}
	}

	if len(l.registerRouters) == 0 {
		l.registerRouters = append(l.registerRouters, httpSelfIntroduction)
	}

	for _, f := range l.registerRouters {
		f(r)
	}

	l.server.Handler = l.engine

	return l
}

// WithServer set http server
func WithServer(s *http.Server) LierGinOption {
	return func(l *LierGin) {
		l.server = s
	}
}

// WithRegisterRouters set register router func
func WithRegisterRouters(f ...Router) LierGinOption {
	return func(l *LierGin) {
		l.registerRouters = f
	}
}

// httpSelfIntroduction 自我介绍
func httpSelfIntroduction(router gin.IRouter) {
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

func RegisterPing(router gin.IRouter) {
	router.GET("/ping", Ping)
}

// Ping the server
func Ping(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusOK)
}

func (l *LierGin) Start() error {
	fmt.Println("GinHttpServer starting...")
	fmt.Println("listen on ", l.server.Addr)
	return l.server.ListenAndServe()
}

// Stop http server stop
func (l *LierGin) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := l.server.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("GinHttpServer stop")
}
