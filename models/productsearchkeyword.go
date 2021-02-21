package models

type ProductSearchKeyword struct {
	Entity
	LanguageID string    `json:"languageId"`
	ProductID  string    `json:"productId"`
	Keyword    string    `json:"keyword"`
	Ranking    float64   `json:"ranking"`
	Product    *Product  `json:"product"`
	Language   *Language `json:"language"`
}
