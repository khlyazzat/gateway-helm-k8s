package middleware

import (
	"strconv"
	"time"

	"main/services/auth-service/internal/metrics"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(c *gin.Context) {
	start := time.Now()
	method := c.Request.Method
	path := c.FullPath()

	if path == "" {
		path = "unknown"
	}

	defer func() {
		duration := time.Since(start).Seconds()
		statusCode := c.Writer.Status()

		metrics.HttpRequests.WithLabelValues(method, path, strconv.Itoa(statusCode)).Inc()
		metrics.HttpDuration.WithLabelValues(method, path).Observe(duration)

		if statusCode >= 500 {
			metrics.HttpErrors.WithLabelValues(method, path).Inc()
		}
	}()

	c.Next()
}
