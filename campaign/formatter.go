package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formattedCampaign := CampaignFormatter{}

	formattedCampaign.ID = campaign.ID
	formattedCampaign.UserID = campaign.UserID
	formattedCampaign.Name = campaign.Name
	formattedCampaign.ShortDescription = campaign.ShortDescription
	formattedCampaign.GoalAmount = campaign.GoalAmount
	formattedCampaign.CurrentAmount = campaign.CurrentAmount
	formattedCampaign.Slug = campaign.Slug
	formattedCampaign.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		formattedCampaign.ImageURL = campaign.CampaignImages[0].FileName
	}

	return formattedCampaign
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	formattedCampaigns := []CampaignFormatter{}

	for _, campaign := range campaigns {
		formattedCampaign := FormatCampaign(campaign)
		formattedCampaigns = append(formattedCampaigns, formattedCampaign)
	}

	return formattedCampaigns
}
