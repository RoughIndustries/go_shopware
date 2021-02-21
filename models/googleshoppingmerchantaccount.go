package models

type GoogleShoppingMerchantAccount struct {
	Entity
	AccountID  string                 `json:"accountId"`
	MerchantID string                 `json:"merchantId"`
	Account    *GoogleShoppingAccount `json:"account"`
}
