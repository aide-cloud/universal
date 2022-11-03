package p8s

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var responseGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "api_response_total",
	},
	[]string{"method", "path", "status"},
)

func init() {
	prometheus.MustRegister(responseGauge)
}

// ResponseTime 响应时间统计中间件
func ResponseTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(begin)
		responseGauge.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprintf("%d", c.Writer.Status())).Set(latency.Seconds())
	}
}
