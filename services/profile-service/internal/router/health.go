package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthClient struct{}

func NewHealthClient() Router {
	return &healthClient{}
}

func (c *healthClient) RegisterAdminRouter(_ *gin.RouterGroup) {}

func (c *healthClient) RegisterRouter(g *gin.RouterGroup) {
	healthGroup := g.Group("")
	healthGroup.GET("/health", c.HealthCheck)
}

func (c *healthClient) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
