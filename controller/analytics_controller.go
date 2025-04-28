package controller

import (
	"ad-server/services"
	"ad-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnalyticsController struct {
	Service services.IAnalyticsService
}

func NewAnalyticsController(service services.IAnalyticsService) *AnalyticsController {
	return &AnalyticsController{Service: service}
}

func (c *AnalyticsController) GetAdClickCounts(ctx *gin.Context) {
	clickCounts, err := c.Service.GetAdClickCounts()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	utils.SuccessResponse(ctx, clickCounts)
}

func (c *AnalyticsController) GetHourlyAnalytics(ctx *gin.Context) {
	hourlyData, err := c.Service.GetHourlyAnalytics()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	utils.SuccessResponse(ctx, hourlyData)
}
