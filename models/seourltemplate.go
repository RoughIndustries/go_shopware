package models

type SeoURLTemplate struct {
	Entity
	SalesChannelID string        `json:"salesChannelId"`
	EntityName     string        `json:"entityName"`
	RouteName      string        `json:"routeName"`
	Template       string        `json:"template"`
	IsValid        string        `json:"isValid"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
}
