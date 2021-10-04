package campaign

import "strings"

type CampaignFormatter struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	Slug string `json:"slug"`

}

func FormatCampaign(campaign Campaign) CampaignFormatter  {
	campaignFormatter := CampaignFormatter{}

	campaignFormatter.Id = campaign.Id

	campaignFormatter.UserId = campaign.UserId

	campaignFormatter.Name = campaign.Name

	campaignFormatter.ShortDescription = campaign.ShortDescription

	campaignFormatter.GoalAmount = campaign.GoalAmount

	campaignFormatter.CurrentAmount = campaign.CurrentAmount

	campaignFormatter.Slug = campaign.Slug

	campaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter  {

	campaignsFormat :=  []CampaignFormatter{}

	for _, v := range campaigns {
		campaign := FormatCampaign(v)
		campaignsFormat = append(campaignsFormat, campaign)
	}

	return campaignsFormat
}

type CampaignDetailFormatter struct{
	Id int `json:"id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description string `json:"description"`
	ImageUrl string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	UserId int `json:"user_id"`
	Slug string `json:"slug"`
	Perks []string `json:"perks"`
	User CampaignUserFormatter `json:"user"`
	Images []CampaignImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name string `json:"name"`
	ImageUrl string `json:"image_url"` 
}

type CampaignImagesFormatter struct {
	ImageUrl string `json:"image_url"`
	IsPrimary bool `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter  {
	campaignDetailFormat := CampaignDetailFormatter{}

	campaignDetailFormat.Id = campaign.Id

	campaignDetailFormat.Name = campaign.Name

	campaignDetailFormat.ShortDescription = campaign.ShortDescription

	campaignDetailFormat.Description = campaign.Description

	campaignDetailFormat.ImageUrl = ""

	campaignDetailFormat.GoalAmount = campaign.GoalAmount

	campaignDetailFormat.CurrentAmount = campaign.CurrentAmount

	campaignDetailFormat.UserId = campaign.UserId

	campaignDetailFormat.Slug = campaign.Slug

	

	if len(campaign.CampaignImages)  > 0 {
		campaignDetailFormat.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, v := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(v))
	}

	// perks = append(perks,strings.Split(campaign.Perks,","))

	campaignDetailFormat.Perks = perks
	
	user := campaign.User

	campaignUserFormatter := CampaignUserFormatter{}

	campaignUserFormatter.Name = user.Name

	campaignUserFormatter.ImageUrl = user.AvatarFileName

	campaignDetailFormat.User = campaignUserFormatter

	images := []CampaignImagesFormatter{}


	for _,v := range  campaign.CampaignImages {
		imageFormat := CampaignImagesFormatter{}
		imageFormat.ImageUrl = v.FileName

		isPrimary := false

		if v.IsPrimary == 1 {
			isPrimary = true
		}

		imageFormat.IsPrimary = isPrimary

		images = append(images, imageFormat)
	}

	campaignDetailFormat.Images = images

	



	return campaignDetailFormat

}