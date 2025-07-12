package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	RegisterRouter(*gin.RouterGroup)
	RegisterAdminRouter(*gin.RouterGroup)
}
