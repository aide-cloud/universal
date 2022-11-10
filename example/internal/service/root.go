package service

import (
	"context"
	"embed"
	"fmt"
	"github.com/aide-cloud/universal/example/internal/conf"
	"github.com/aide-cloud/universal/graphql"
	"github.com/gin-gonic/gin"
)

type Root struct{}

// Content holds all the SDL file content.
//go:embed sdl
var content embed.FS

// Ping is a ping service.
func (r *Root) Ping() *string {
	res := "pong"
	return &res
}

func (r *Root) PingString(_ context.Context, args struct {
	Ping string
}) string {
	return fmt.Sprintf("Hello, %s", args.Ping)
}

func RegisterRoot(r *gin.Engine) {
	isDev := conf.GetConfig().Server.Mode != gin.ReleaseMode
	graphql.RegisterHttpRouter(r, new(Root), content, isDev)
}
