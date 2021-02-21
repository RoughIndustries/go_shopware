package models

type CountryStateTranslation struct {
	TranslationEntity
	CountryStateID string        `json:"countryStateId"`
	Name           string        `json:"name"`
	CountryState   *CountryState `json:"countryState"`
}
