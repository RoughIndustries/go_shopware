package client

import (
	"github.com/RoughIndustries/go_shopware/models"
	"net/http"
)

func (c *Client) GetASingleProduct(productId string) (*models.ProductWrapper, error) {
	product := &models.ProductWrapper{}
	err := c.do(http.MethodPost, "/store-api/v3/product/"+productId, nil, &product)
	return product, err
}
