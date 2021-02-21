package models

type UnitTranslation struct {
	TranslationEntity
	UnitID    string `json:"unitId"`
	ShortCode string `json:"shortCode"`
	Name      string `json:"name"`
	Unit      *Unit  `json:"unit"`
}
