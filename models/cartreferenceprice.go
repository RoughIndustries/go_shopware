package models

type ReferencePrice struct {
	ReferencePriceDefinition
	Price float64 `json:"price"`
}

func NewReferencePrice(price, purchaseUnit, referenceUnit float64, unitName string) *ReferencePrice {
	return &ReferencePrice{
		Price: price,
		ReferencePriceDefinition: ReferencePriceDefinition{
			PurchaseUnit:  purchaseUnit,
			ReferenceUnit: referenceUnit,
			UnitName:      unitName,
		},
	}
}

func (r *ReferencePrice) GetPrice() float64 {
	return r.Price
}

func (r *ReferencePrice) GetAPIAlias() string {
	return "cart_price_reference"
}
