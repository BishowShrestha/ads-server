package controller

import (
	"ad-server/model"
	"ad-server/services"
	"ad-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClicksController struct {
	ClickService services.IClicksService
}

func NewClicksController(clickService services.IClicksService) *ClicksController {
	return &ClicksController{ClickService: clickService}
}

func (c *ClicksController) SaveClick(ctx *gin.Context) {
	var clickEvent model.ClickEvent
	if err := ctx.ShouldBindJSON(&clickEvent); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input")
		return
	}
	err := c.ClickService.SaveClick(clickEvent)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to save click")
		return
	}
	utils.SuccessResponse(ctx, gin.H{"message": "Click saved successfully"})
}
