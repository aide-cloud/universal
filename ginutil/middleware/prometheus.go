package middleware

import (
	"fmt"
	"github.com/aide-cloud/universal/alog"
	"github.com/aide-cloud/universal/helper/runtimehelper"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	requestGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "api_request_total",
		},
		[]string{"method", "path", "status"},
	)

	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_request_total",
		},
		[]string{"method", "path", "status"},
	)
)

func init() {
	prometheus.MustRegister(requestGauge, requestCounter)
}

// Prometheus 响应时间统计中间件
func Prometheus(logger alog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		c.Next()
		go func() {
			runtimehelper.Recover("ResponseTime panic", runtimehelper.NewRecoverConfig(runtimehelper.WithLog(logger)))
			end := time.Now()
			latency := end.Sub(begin)
			requestGauge.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprintf("%d", c.Writer.Status())).Set(float64(latency.Nanoseconds()))
		}()

		go func() {
			runtimehelper.Recover("ResponseTime panic", runtimehelper.NewRecoverConfig(runtimehelper.WithLog(logger)))
			requestCounter.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprintf("%d", c.Writer.Status())).Inc()
		}()
	}
}
