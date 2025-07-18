package router

import (
	"main/services/api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	RegisterRouter(*gin.RouterGroup, middleware.Middleware)
	RegisterAdminRouter(*gin.RouterGroup, middleware.Middleware)
}
