package proxy

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyToAuth(c *gin.Context) {
	proxyRequest(c, "http://auth-service:8080")
}

func ProxyToProfile(c *gin.Context) {
	proxyRequest(c, "http://profile-service:8080")
}

func proxyRequest(c *gin.Context, targetHost string) {
	targetURL := targetHost + c.Request.URL.Path

	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
		return
	}

	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to proxy request"})
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}
	c.Writer.WriteHeader(resp.StatusCode)

	bodyBytes, _ := io.ReadAll(resp.Body)
	c.Writer.Write(bodyBytes)
}
