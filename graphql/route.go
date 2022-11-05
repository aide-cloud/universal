package graphql

import (
	"embed"
	"github.com/aide-cloud/universal/basic/assert"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

func NewGraphQLHandlerFunc() http.HandlerFunc {
	var page = []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
		</head>
		<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function graphQLFetcher(graphQLParams) {
					return fetch("/graphql", {
						method: "post",
						body: JSON.stringify(graphQLParams),
						credentials: "include",
					}).then(function (response) {
						return response.text();
					}).then(function (responseBody) {
						try {
							return JSON.parse(responseBody);
						} catch (error) {
							return responseBody;
						}
					});
				}

				ReactDOM.render(
					React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
					document.getElementById("graphiql")
				);
			</script>
		</body>
	</html>
	`)

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(page)
		if err != nil {
			log.Printf("write page error: %s\n", err)
			return
		}
	}
}

func NewHandler(root any, content embed.FS) *relay.Handler {
	// 判断root是否为结构体指针或者结构体指针
	if !assert.IsStruct(root) && !assert.IsStructPtr(root) {
		log.Fatal("root must be a struct pointer")
	}
	s, err := String(content)
	if err != nil {
		log.Fatalf("reading embedded schema contents: %s", err)
	}

	return &relay.Handler{Schema: graphql.MustParseSchema(s, root)}
}

func RegisterHttpRouter(r *gin.Engine, root any, content embed.FS, isDev ...bool) {
	if len(isDev) > 0 && isDev[0] {
		r.GET("/graphql", gin.WrapF(NewGraphQLHandlerFunc()))
	}
	r.POST("/graphql", gin.WrapH(NewHandler(root, content)))
}
