package routes

import (
	"embed"
	"github.com/aide-cloud/universal/graphql"
	"github.com/gin-gonic/gin"
)

func RegisterHttpRouter(r *gin.Engine, root any, content embed.FS, isDev ...bool) {
	if len(isDev) > 0 && isDev[0] {
		r.GET("/graphql", gin.WrapF(graphql.NewGraphQLHandlerFunc()))
	}
	r.POST("/graphql", gin.WrapH(graphql.NewHandler(root, content)))
}

// GinGraphQLHandlerFunc returns a http.HandlerFunc that can be used to serve the GraphiQL IDE.
func GinGraphQLHandlerFunc() gin.HandlerFunc {
	return gin.WrapF(graphql.NewGraphQLHandlerFunc())
}

// GinHandler returns a http.Handler that can be used to serve the GraphQL API.
func GinHandler(root any, content embed.FS) gin.HandlerFunc {
	return gin.WrapH(graphql.NewHandler(root, content))
}
