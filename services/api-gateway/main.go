package main

import (
	"context"
	"log"
	"net/http"
	"time"

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

	srv := &http.Server{
		Addr:    cfg.HTTPConfig.Port,
		Handler: router,
	}

	ctx, cancel := utils.GracefulShutdown(context.Background())
	defer cancel()

	go func() {
		log.Println("Starting server on", cfg.HTTPConfig.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Graceful shutdown failed: %s", err)
	}

	log.Println("Server stopped")
}
