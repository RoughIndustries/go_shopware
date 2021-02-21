package models

type SalesChannelType struct {
	Entity
	Name            string                         `json:"name"`
	Manufacturer    string                         `json:"manufacturer"`
	Description     string                         `json:"description"`
	DescriptionLong string                         `json:"descriptionLong"`
	CoverURL        string                         `json:"coverUrl"`
	IconName        string                         `json:"iconName"`
	ScreenshotUrls  []string                       `json:"screenshotUrls"`
	SalesChannels   []*SalesChannel                `json:"salesChannels"`
	translations    []*SalesChannelTypeTranslation `json:"translations"`
}
