package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"main/pkg/jwt"
	"main/services/auth-service/internal/config"
	"main/services/auth-service/internal/db/cache"
	"main/services/auth-service/internal/db/postgres"
	"main/services/auth-service/internal/metrics"
	"main/services/auth-service/internal/middleware"
	apiRouter "main/services/auth-service/internal/router"
	"main/utils"

	"main/services/auth-service/internal/auth"
)

func main() {
	metrics.Init()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	db := postgres.New(cfg.DBConfig)

	cacheClient, err := cache.NewCache(cfg.RedisConfig)
	if err != nil {
		panic(err)
	}

	tokenCache := cache.NewTokenCache(cacheClient, cfg.JwtConfig.RefreshTTL)

	jwtClient := jwt.New(jwt.Config{
		APISecret:     cfg.JwtConfig.APISecret,
		RefreshSecret: cfg.JwtConfig.RefreshSecret,
		AccessTTL:     cfg.JwtConfig.AccessTTL,
		RefreshTTL:    cfg.JwtConfig.RefreshTTL,
	})

	router := gin.New()

	router.GET("/metrics", metrics.Handler())
	router.Use(middleware.MetricsMiddleware)

	v1 := router.Group("/v1")

	healthCClient := apiRouter.NewHealthClient()
	healthCClient.RegisterRouter(v1)

	authService := auth.NewAuthService(db, tokenCache, jwtClient)
	authClient := apiRouter.NewAuthClient(authService)
	authClient.RegisterRouter(v1)

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
