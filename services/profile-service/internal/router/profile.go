package router

import (
	"main/services/profile-service/internal/dto"
	"main/services/profile-service/internal/profile"
	"net/http"

	"github.com/gin-gonic/gin"
)

type profileClient struct {
	service profile.Profile
}

func NewProfileClient(a profile.Profile) Router {
	return &profileClient{
		service: a,
	}
}

func (c *profileClient) RegisterRouter(g *gin.RouterGroup) {
	profileGroup := g.Group("/profile")
	profileGroup.GET("/get", c.GetProfile)
	profileGroup.PUT("/update", c.UpdateProfile)
}

func (c *profileClient) RegisterAdminRouter(_ *gin.RouterGroup) {}

func (c *profileClient) GetProfile(ctx *gin.Context) {
	userIDContext, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	userID, ok := userIDContext.(int64)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid userID type"})
		return
	}
	userEmailContext, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	email, ok := userEmailContext.(string)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid email type"})
		return
	}
	resp, err := c.service.GetProfile(ctx, userID, email)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *profileClient) UpdateProfile(ctx *gin.Context) {
	var body dto.UpdateProfileRequest

	userIDContext, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	userID, ok := userIDContext.(int64)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid userID type"})
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		_ = ctx.Error(err)
		return
	}
	err := c.service.UpdateProfile(ctx, userID, &body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "profile updated"})
}
