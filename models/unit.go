package models

type Unit struct {
	Entity
	ShortCode    string             `json:"shortCode"`
	Name         string             `json:"name"`
	Translations []*UnitTranslation `json:"translations"`
	Products     []*Product         `json:"products"`
}
