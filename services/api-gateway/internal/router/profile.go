package router

import (
	"main/services/api-gateway/internal/middleware"
	"main/services/api-gateway/internal/proxy"

	"github.com/gin-gonic/gin"
)

type profileClient struct {
}

func (c *profileClient) RegisterRouter(g *gin.RouterGroup, _ middleware.Middleware) {
	profileGroup := g.Group("/profile")
	profileGroup.GET("/get", proxy.ProxyToProfile)
	profileGroup.PUT("/update", proxy.ProxyToProfile)
}

func (c *profileClient) RegisterAdminRouter(_ *gin.RouterGroup, _ middleware.Middleware) {}

func NewProfileClient() Router {
	return &profileClient{}
}
