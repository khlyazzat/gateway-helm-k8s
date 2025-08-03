package proxy

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyToAuth(ctx *gin.Context) {
	proxyRequest(ctx, "http://auth-service")
}

func ProxyToProfile(ctx *gin.Context) {
	proxyRequest(ctx, "http://profile-service")
}

func proxyRequest(ctx *gin.Context, targetHost string) {
	targetURL := targetHost + ctx.Request.URL.Path

	req, err := http.NewRequest(ctx.Request.Method, targetURL, ctx.Request.Body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
		return
	}

	if targetHost == "http://profile-service" {
		email := ctx.GetString("email")

		req.Header.Add("X-User-Email", email)
	}

	for k, v := range ctx.Request.Header {
		req.Header[k] = v
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "failed to proxy request"})
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		ctx.Writer.Header()[k] = v
	}
	ctx.Writer.WriteHeader(resp.StatusCode)

	bodyBytes, _ := io.ReadAll(resp.Body)
	ctx.Writer.Write(bodyBytes)
}
