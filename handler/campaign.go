package handler

import (
	"api-crowdfunding/campaign"
	"api-crowdfunding/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// tangkap parameter di handler
// handler ke service
// service yang menentukan repository mana yang di call
// repository : GetAll, GetByUserID
// db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns (c *gin.Context) {
	userID,_ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error To Get Campaign", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List Of Campaign", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
	return
}

