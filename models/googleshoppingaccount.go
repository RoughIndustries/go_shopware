package models

type GoogleShoppingAccount struct {
	Entity
	SalesChannelID                string                         `json:"salesChannelId"`
	Email                         string                         `json:"email"`
	name                          string                         `json:"name"`
	credential                    *GoogleAccountCredential       `json:"credential"`
	salesChannel                  *SalesChannel                  `json:"salesChannel"`
	googleShoppingMerchantAccount *GoogleShoppingMerchantAccount `json:"googleShoppingMerchantAccount"`
}
