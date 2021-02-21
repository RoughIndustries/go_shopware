package models

type Locale struct {
	Entity
	Code         string               `json:"code"`
	Name         string               `json:"name"`
	Territory    string               `json:"territory"`
	Translations []*LocaleTranslation `json:"translations"`
	Users        []*User              `json:"users"`
	Languages    []*Language          `json:"languages"`
}
