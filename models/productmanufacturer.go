package models

type ProductManufacturer struct {
	Entity
	Description  string                            `json:"description"`
	Link         string                            `json:"link"`
	Media        MediaEntity                       `json:"media"`
	MediaID      string                            `json:"mediaId"`
	Name         string                            `json:"name"`
	Products     []*Product                        `json:"products"`
	Translations []*ProductManufacturerTranslation `json:"translations"`
}

type ProductManufacturerTranslation struct {
	TranslationEntity
	ProductManufacturerID string               `json:"productManufacturerId"`
	Name                  string               `json:"name"`
	Description           string               `json:"description"`
	ProductManufacturer   *ProductManufacturer `json:"productManufacturer"`
}
