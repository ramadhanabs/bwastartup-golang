package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	fmt.Println(err)

	campaigns, err := h.service.FindCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns), nil)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	input := campaign.GetCampaignDetailInput{}

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	detailCampaign, err := h.service.FindCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", campaign.FormatDetailCampaign(detailCampaign), nil)
	c.JSON(http.StatusOK, response)
}
