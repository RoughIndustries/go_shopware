package models

type CountryState struct {
	Entity
	CountryID         string                     `json:"countryId"`
	ShortCode         string                     `json:"shortCode"`
	Name              string                     `json:"name"`
	Position          int                        `json:"position"`
	Active            bool                       `json:"active"`
	Country           *Country                   `json:"country"`
	Translations      []*CountryStateTranslation `json:"translations"`
	CustomerAddresses []*CustomerAddress         `json:"customerAddresses"`
	OrderAddresses    []*OrderAddress            `json:"orderAddresses"`
}
