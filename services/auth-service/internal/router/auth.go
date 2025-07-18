package router

import (
	"errors"
	"main/services/auth-service/internal/auth"
	"main/services/auth-service/internal/dto"
	"main/services/auth-service/internal/values"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authClient struct {
	service auth.Auth
}

func NewAuthClient(a auth.Auth) Router {
	return &authClient{
		service: a,
	}
}

func (c *authClient) RegisterRouter(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	authGroup.POST("/sign-up", c.SignUp)
	authGroup.POST("/sign-in", c.SignIp)
	authGroup.POST("/refresh", c.Refresh)
}

func (c *authClient) RegisterAdminRouter(_ *gin.RouterGroup) {}

func (c *authClient) SignUp(ctx *gin.Context) {
	var body dto.SignUpRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	_, err := c.service.SignUp(ctx, &body)
	if errors.Is(err, values.ErrEmailExists) {
		ctx.JSON(http.StatusConflict, gin.H{"message": "email already exists"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "new user created"})
}

func (c *authClient) SignIp(ctx *gin.Context) {
}

func (c *authClient) Refresh(ctx *gin.Context) {
}
