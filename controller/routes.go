package controller

import (
	"ad-server/repository"
	"ad-server/services"
)

func (ctl *Controller) Routes() {
	adRepository := repository.NewAdsRepository(ctl.DB)
	adService := services.NewAdsService(adRepository)
	adController := NewAdsController(adService)

	clicksRepository := repository.NewClicksRepository(ctl.DB, ctl.Logger)
	clicksService := services.NewClicksService(clicksRepository)
	clicksController := NewClicksController(clicksService)

	analyticsService := services.NewAnalyticsService(clicksRepository)
	analyticsController := NewAnalyticsController(analyticsService)
	adsGroup := ctl.Gin.Group("/ads")
	{
		adsGroup.POST("/", adController.CreateAds)
		adsGroup.POST("/click", clicksController.SaveClick)
		adsGroup.GET("", adController.GetAllAds)
		adsGroup.GET("analytics", analyticsController.GetAdClickCounts)
		adsGroup.GET("analytics/hourly", analyticsController.GetHourlyAnalytics)
	}
}
