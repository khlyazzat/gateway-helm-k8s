package router

import (
	"main/services/api-gateway/internal/middleware"
	"main/services/api-gateway/internal/proxy"

	"github.com/gin-gonic/gin"
)

type authClient struct {
}

func (c *authClient) RegisterRouter(g *gin.RouterGroup, _ middleware.Middleware) {
	authGroup := g.Group("/auth")
	authGroup.POST("/sign-up", proxy.ProxyToAuth)
	authGroup.POST("/sign-in", proxy.ProxyToAuth)
	authGroup.POST("/refresh", proxy.ProxyToAuth)
}

func (c *authClient) RegisterAdminRouter(_ *gin.RouterGroup, _ middleware.Middleware) {}

func NewAuthClient() Router {
	return &authClient{}
}
