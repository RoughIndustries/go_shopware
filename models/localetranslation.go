package models

type LocaleTranslation struct {
	TranslationEntity
	LocaleID  string  `json:"localeId"`
	Name      string  `json:"name"`
	Territory string  `json:"territory"`
	Locale    *Locale `json:"locale"`
}
