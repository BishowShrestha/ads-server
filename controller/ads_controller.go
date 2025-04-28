package controller

import (
	"ad-server/model"
	"ad-server/services"
	"ad-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdsController struct {
	Service services.IAdsService
}

func NewAdsController(service services.IAdsService) *AdsController {
	return &AdsController{Service: service}
}

func (c *AdsController) CreateAds(ctx *gin.Context) {
	var ad model.Ad
	if err := ctx.ShouldBindJSON(&ad); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input")
		return
	}
	err := c.Service.CreateAds(ad)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create ad")
		return
	}
	utils.SuccessResponse(ctx, gin.H{"message": "Ad created successfully"})
}

func (c *AdsController) GetAllAds(ctx *gin.Context) {
	ads, err := c.Service.GetAllAds()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(ctx, ads)
}
