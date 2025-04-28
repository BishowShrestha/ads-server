package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ClickCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "click_events_total",
			Help: "Total click events received",
		})
)

func RegisterMetrics(r *gin.Engine) {
	prometheus.MustRegister(ClickCounter)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
