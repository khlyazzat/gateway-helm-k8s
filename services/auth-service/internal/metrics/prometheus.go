package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	HttpRequests *prometheus.CounterVec
	HttpDuration *prometheus.HistogramVec
	HttpErrors   *prometheus.CounterVec
)

func Init() {
	HttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	HttpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of latencies",
			Buckets: prometheus.ExponentialBuckets(0.01, 2, 15),
		},
		[]string{"method", "path"},
	)

	HttpErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of 5xx errors",
		},
		[]string{"method", "path"},
	)

	prometheus.MustRegister(HttpRequests, HttpDuration, HttpErrors)
}

func Handler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
