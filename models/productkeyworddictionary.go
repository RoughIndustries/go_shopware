package models

type ProductKeywordDictionary struct {
	Entity
	ID         string    `json:"id"`
	LanguageID string    `json:"languageId"`
	Keyword    string    `json:"keyword"`
	Reversed   string    `json:"reversed"`
	Language   *Language `json:"language"`
}
