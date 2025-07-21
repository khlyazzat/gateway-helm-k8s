package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	apiRouter "main/services/api-gateway/internal/router"

	"main/pkg/jwt"
	"main/services/api-gateway/internal/config"
	"main/services/api-gateway/internal/middleware"
	"main/utils"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		log.Println("Incoming request path:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	jwtClient := jwt.New(jwt.Config{
		// Issuer:        cfg.AppConfig.AppName,
		APISecret:     cfg.JwtConfig.APISecret,
		RefreshSecret: cfg.JwtConfig.RefreshSecret,
		AccessTTL:     cfg.JwtConfig.AccessTTL,
		RefreshTTL:    cfg.JwtConfig.RefreshTTL,
	})

	mwService := middleware.New(jwtClient)

	v1 := router.Group("/v1", mwService.Error, mwService.Panic)
	authorize := v1.Group("", mwService.Authorize)

	authClient := apiRouter.NewAuthClient()
	authClient.RegisterRouter(v1, mwService)

	profileClient := apiRouter.NewProfileClient()
	profileClient.RegisterRouter(authorize, mwService)

	// for _, ri := range router.Routes() {
	// 	log.Printf("Route registered: %s %s\n", ri.Method, ri.Path)
	// }

	srv := &http.Server{
		Addr:    cfg.HTTPConfig.Port,
		Handler: router,
	}

	go func() {
		log.Println("Starting server on", cfg.HTTPConfig.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Web server error: %s", err)
		}
	}()

	ctx, cancel := utils.GracefulShutdown(context.TODO())
	defer cancel()

	<-ctx.Done()
}
