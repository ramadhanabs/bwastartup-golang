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
	fmt.Printf("total campaigns %v", len(campaigns))
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns), nil)
	c.JSON(http.StatusOK, response)
}
