package main

import (
	"context"
	"log"
	"main/services/profile-service/internal/config"
	"main/services/profile-service/internal/db/postgres"
	"main/services/profile-service/internal/metrics"
	"main/services/profile-service/internal/middleware"
	"main/services/profile-service/internal/profile"
	"main/utils"
	"net/http"

	apiRouter "main/services/profile-service/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	metrics.Init()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	db := postgres.New(cfg.DBConfig)

	router := gin.New()

	router.GET("/metrics", metrics.Handler())
	router.Use(middleware.MetricsMiddleware)

	router.Use(middleware.HeadersMiddleware())

	// router.Use(func(c *gin.Context) {
	// 	log.Println("Incoming request path:", c.Request.Method, c.Request.URL.Path)
	// 	c.Next()
	// })

	v1 := router.Group("/v1")

	healthCClient := apiRouter.NewHealthClient()
	healthCClient.RegisterRouter(v1)

	profileService := profile.NewProfileService(db)
	profileClient := apiRouter.NewProfileClient(profileService)
	profileClient.RegisterRouter(v1)

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
