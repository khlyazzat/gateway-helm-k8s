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
		_ = ctx.Error(err)
		return
	}
	_, err := c.service.SignUp(ctx, &body)
	if errors.Is(err, values.ErrEmailExists) {
		_ = ctx.Error(err)
		return
	}
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "new user created"}) //TODO
}

func (c *authClient) SignIp(ctx *gin.Context) {
	var body dto.SignInRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		_ = ctx.Error(err)
		return
	}
	resp, err := c.service.SignIn(ctx, &body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *authClient) Refresh(ctx *gin.Context) {
	var body dto.RefreshRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		_ = ctx.Error(err)
		return
	}
	resp, err := c.service.Refresh(ctx, &body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
