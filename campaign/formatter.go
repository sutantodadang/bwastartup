package campaign

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