package models

type CustomerGroup struct {
	Entity
	Name          string                      `json:"name"`
	DisplayGross  bool                        `json:"displayGross"`
	Translations  []*CustomerGroupTranslation `json:"translations"`
	Customers     []*Customer                 `json:"customers"`
	SalesChannels []*SalesChannel             `json:"salesChannels"`
}
