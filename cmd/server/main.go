package main

import (
	"ad-server/internal/ads"
	"ad-server/internal/analytics"
	"ad-server/internal/clicks"
	"ad-server/internal/db"
	"ad-server/pkg/config"
	"ad-server/pkg/logger"
	"ad-server/pkg/metrics"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.NewLogger()
	dbConn := db.NewPostgresDB(cfg.DatabaseURL)

	clickRepo := clicks.NewClickRepository(dbConn, log)
	adRepo := ads.NewAdRepository(dbConn, log)

	clickService := clicks.NewClickService(clickRepo)
	adService := ads.NewAdService(adRepo)
	analyticsService := analytics.NewAnalyticsService(clickRepo)

	clickWorker := clicks.NewClickWorker(clickService)
	go clickWorker.Start()

	r := gin.Default()

	metrics.RegisterMetrics(r)

	ads.RegisterRoutes(r, adService)
	clicks.RegisterRoutes(r, clickWorker)
	analytics.RegisterRoutes(r, analyticsService)

	log.Info("Server starting", zap.String("port", cfg.Port))
	r.Run(":" + cfg.Port)
}
