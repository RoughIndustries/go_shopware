package models

type CategoryTranslation struct {
	TranslationEntity
	CategoryID      string    `json:"categoryId"`
	Name            string    `json:"name"`
	Breadcrumb      []string  `json:"breadcrumb"`
	Category        *Category `json:"category"`
	Language        *Language `json:"language"`
	SlotConfig      []string  `json:"slotConfig"`
	ExternalLink    string    `json:"externalLink"`
	Description     string    `json:"description"`
	MetaTitle       string    `json:"metaTitle"`
	MetaDescription string    `json:"metaDescription"`
	Keywords        string    `json:"keywords"`
}
