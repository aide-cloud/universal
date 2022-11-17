package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterMetricRoute registers the metric route.
func RegisterMetricRoute(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
