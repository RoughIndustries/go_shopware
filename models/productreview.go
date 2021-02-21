package models

type ProductReview struct {
	Entity
	ProductID      string        `json:"productID"`
	CustomerID     string        `json:"customerId"`
	SalesChannelID string        `json:"salesChannelId"`
	LanguageID     string        `json:"languageId"`
	ExternalUser   string        `json:"externalUser"`
	ExternalEmail  string        `json:"externalEmail"`
	Points         float64       `json:"points"`
	Status         bool          `json:"status"`
	Comment        string        `json:"comment"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
	Language       *Language     `json:"language"`
	Customer       *Customer     `json:"customer"`
	Product        *Product      `json:"product"`
	Content        string        `json:"content"`
	Title          string        `json:"title"`
}
