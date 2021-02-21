package models

type ProductMediaEntity struct {
	Entity
	Media     *MediaEntity `json:"media"`
	MediaID   string       `json:"mediaId"`
	Position  int          `json:"position"`
	Product   *Product     `json:"product"`
	ProductID string       `json:"productId"`
}
