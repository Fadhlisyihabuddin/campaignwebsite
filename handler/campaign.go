package handler

import (
	"campaignwebsite/campaign"
	"campaignwebsite/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) FindCampaigns(c *gin.Context) {

	UserID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaign(UserID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(400, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)

	response := helper.ApiResponse("Succes to get Campaign", http.StatusOK, "succes", formatter)
	c.JSON(200, response)
}

func (h *campaignHandler) GetCampaignDetail(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	if err := c.ShouldBindUri(&input); err != nil {
		response := helper.ApiResponse(" Failed to get Campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaigns, err := h.campaignService.GetCampaignDetail(input)
	if err != nil {
		response := helper.ApiResponse(" Failed to get Campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaignDetail(campaigns)

	response := helper.ApiResponse("Login Succes", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)
}
