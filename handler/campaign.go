package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler  {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context)  {
	userId, _ := strconv.Atoi(c.Query("user_id")) 

	campaigns, err := h.service.GetCampaigns(userId)

	if err != nil {

		res := helper.APIResponse("Get Campaign Failed",http.StatusBadRequest,"Error",nil)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.APIResponse("Get Campaign Success",http.StatusOK,"success",campaign.FormatCampaigns(campaigns))

	c.JSON( http.StatusOK,res)
}

func (h *campaignHandler) GetCampaign(c *gin.Context){
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {

		res := helper.APIResponse("Failed Get Detail Campaign", http.StatusBadRequest, "error",nil)
		c.JSON(http.StatusBadRequest,res)
		return
	}

	campaignDetail, err := h.service.GetCampaignById(input)

	if err != nil {

		res := helper.APIResponse("Failed Get Detail Campaign", http.StatusBadRequest, "error",nil)
		c.JSON(http.StatusBadRequest,res)
		return
	}

	res := helper.APIResponse("Campaign Detail Success", http.StatusOK,"success",campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK,res)


}

func (h *campaignHandler) CreateCampaign(c *gin.Context)  {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorValid := helper.ErrorValidationFormat(err)

		errorMsg := gin.H{"errors":errorValid}

		res := helper.APIResponse("Failed Create Campaign",http.StatusUnprocessableEntity,"error",errorMsg)
		c.JSON(http.StatusUnprocessableEntity,res)
		return 
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign , err := h.service.CreateCampaign(input)

	if err != nil {
		
		res := helper.APIResponse("Failed Create Campaign",http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest,res)
		return 
	}

	res := helper.APIResponse("Success Create Campaign",http.StatusCreated,"success",campaign.FormatCampaign(newCampaign))

	c.JSON(http.StatusCreated,res)


}