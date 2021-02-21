package models

type ProductConfiguratorSetting struct {
	Entity
	ProductID string               `json:"productId"`
	OptionID  string               `json:"optionId"`
	MediaID   string               `json:"mediaId"`
	Position  int                  `json:"position"`
	Price     []*Price             `json:"price"`
	Option    *PropertyGroupOption `json:"option"`
	Media     *MediaEntity         `json:"media"`
	Selected  bool                 `json:"selected"`
	Product   *Product             `json:"product"`
}
