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
	LierGin struct {
		server             *http.Server
		registerRouterFunc func(router *gin.Engine)
	}

	Config struct {
		Addr           string
		Port           int
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		MaxHeaderBytes int
	}
)

func NewGin(routerFunc ...func(router *gin.Engine)) *LierGin {
	l := &LierGin{}
	if len(routerFunc) > 0 {
		l.SetRouterFunc(routerFunc[0])
	}
	return l
}

func (l *LierGin) SetServer(c *Config) {
	l.server = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", c.Addr, c.Port),
		ReadTimeout:    c.ReadTimeout,
		WriteTimeout:   c.WriteTimeout,
		MaxHeaderBytes: c.MaxHeaderBytes,
	}
}

func (l *LierGin) SetRouterFunc(routerFunc func(router *gin.Engine)) {
	l.registerRouterFunc = routerFunc
}

func (l *LierGin) initGin() {
	// gin初始化
	router := gin.New()

	// 注册路由
	if l.registerRouterFunc == nil {
		l.SetRouterFunc(func(router *gin.Engine) {
			router.GET("/ping", func(context *gin.Context) {
				context.AbortWithStatus(http.StatusOK)
			})
		})
	}

	if l.server == nil {
		l.SetServer(&Config{
			Addr:           LOCALHOST,
			Port:           PORT8080,
			ReadTimeout:    time.Minute,
			WriteTimeout:   time.Minute,
			MaxHeaderBytes: 1 << 20,
		})
	}

	l.registerRouterFunc(router)
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
