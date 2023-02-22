package main

import (
	"embed"
	"github.com/aide-cloud/universal/graphql"
	"github.com/gin-gonic/gin"
)

// Content holds all the SDL file content.
//go:embed sdl
var content embed.FS

type Root struct{}

func (r *Root) Ping() string {
	return "pong"
}

// RegisterHttpRouter registers the GraphQL API and GraphiQL IDE.
func RegisterHttpRouter(r *gin.Engine, root any, content embed.FS, dev ...bool) {
	if len(dev) > 0 && dev[0] {
		r.GET(graphql.DefaultViewPath, GinGraphqlView())
	}

	r.POST(graphql.DefaultHandlePath, GinGraphqlHandler(root, content))
}

// GinGraphqlView returns a http.HandlerFunc that can be used to serve the GraphiQL IDE.
func GinGraphqlView() gin.HandlerFunc {
	return gin.WrapF(graphql.View(graphql.Post, graphql.DefaultHandlePath))
}

// GinGraphqlHandler returns a http.Handler that can be used to serve the GraphQL API.
func GinGraphqlHandler(root any, content embed.FS) gin.HandlerFunc {
	return gin.WrapH(graphql.Handler(root, content))
}

func main() {
	r := gin.Default()
	RegisterHttpRouter(r, &Root{}, content, true)
	r.Run()
}
