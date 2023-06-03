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

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	UserID           int                      `json:"user_id"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             UserFormatter            `json:"user"`
	Images           []CampaignImageFormatter `json:"campaign_images"`
}

type UserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"url"`
	IsPrimary bool   `json:"is_primary"`
}

/*
Formatter Array of Campaign
*/
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

/*
	Formatter Campaign Detail
*/

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	formattedDetailCampaign := CampaignDetailFormatter{}

	formattedDetailCampaign.ID = campaign.ID
	formattedDetailCampaign.UserID = campaign.UserID
	formattedDetailCampaign.Name = campaign.Name
	formattedDetailCampaign.ShortDescription = campaign.ShortDescription
	formattedDetailCampaign.GoalAmount = campaign.GoalAmount
	formattedDetailCampaign.CurrentAmount = campaign.CurrentAmount
	formattedDetailCampaign.Slug = campaign.Slug

	// mapping user
	user := campaign.User
	userCampaign := UserFormatter{}

	userCampaign.Name = user.Name
	userCampaign.ImageURL = user.AvatarFileName

	formattedDetailCampaign.User = userCampaign

	// mapping images
	images := campaign.CampaignImages
	campaignImages := []CampaignImageFormatter{}

	for _, image := range images {
		formattedImage := CampaignImageFormatter{}
		formattedImage.ImageURL = image.FileName
		formattedImage.IsPrimary = image.IsPrimary

		campaignImages = append(campaignImages, formattedImage)
	}

	formattedDetailCampaign.Images = campaignImages

	return formattedDetailCampaign
}
