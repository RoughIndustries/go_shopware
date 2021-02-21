package models

type CountryTranslation struct {
	TranslationEntity
	CountryID string   `json:"countryId"`
	Name      string   `json:"name"`
	Country   *Country `json:"country"`
}
