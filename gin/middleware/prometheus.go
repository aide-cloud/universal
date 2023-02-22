package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	requestGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "api_request",
		},
		[]string{"method", "path", "status"},
	)

	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_request",
		},
		[]string{"method", "path", "status"},
	)
)

func init() {
	prometheus.MustRegister(requestGauge, requestCounter)
}

// Prometheus 响应时间统计中间件
func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(begin)
		requestGauge.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprintf("%d", c.Writer.Status())).Set(float64(latency.Nanoseconds()))
		requestCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprintf("%d", c.Writer.Status())).Inc()
	}
}
