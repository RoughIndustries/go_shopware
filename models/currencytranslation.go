package models

type CurrencyTranslation struct {
	TranslationEntity
	CurrencyID string    `json:"currencyId"`
	ShortName  string    `json:"shortName"`
	Name       string    `json:"name"`
	Currency   *Currency `json:"currency"`
}
