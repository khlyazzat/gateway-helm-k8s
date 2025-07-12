package router

import (
	"main/services/auth-service/internal/auth"

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
	// var body dto.AddUserRequest
	// if err := ctx.ShouldBindJSON(&body); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
	// 	return
	// }
	// res, err := c.service.AddUser(ctx, &body)
	// if errors.Is(err, values.ErrEmailExists) {
	// 	ctx.JSON(http.StatusConflict, gin.H{"message": "email already exists"})
	// 	return
	// }
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
	// 	return
	// }
	// ctx.JSON(http.StatusCreated, gin.H{"id": res.ID})
}

func (c *authClient) SignIp(ctx *gin.Context) {
}

func (c *authClient) Refresh(ctx *gin.Context) {
}
