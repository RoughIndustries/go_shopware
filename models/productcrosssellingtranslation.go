package models

type ProductCrossSellingTranslation struct {
	TranslationEntity
	ProductCrossSellingID string               `json:"productCrossSellingId"`
	Name                  string               `json:"name"`
	ProductCrossSelling   *ProductCrossSelling `json:"productCrossSelling"`
}
