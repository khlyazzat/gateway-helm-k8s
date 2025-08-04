package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/v1/health" {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			return
		}

		userIdHeader := c.Request.Header.Get("X-User-ID")
		if userIdHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		userEmailHeader := c.Request.Header.Get("X-User-Email")
		if userEmailHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
			return
		}
		userID, err := strconv.ParseInt(userIdHeader, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		c.Set("userID", userID)
		c.Set("email", userEmailHeader)

		c.Next()
	}
}
