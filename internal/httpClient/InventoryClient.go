package httpClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/internal/service/dto"
)

type InventoryClient interface {
	CheckAvailability(sku string, quantity int) (bool, error)
}

type InventoryClientImpl struct {
	url        string
	httpClient *http.Client
}

func NewInventoryClient(url string, httpClient *http.Client) InventoryClient {
	return &InventoryClientImpl{
		url:        url,
		httpClient: httpClient,
	}
}

func (c *InventoryClientImpl) CheckAvailability(sku string, quantity int) (bool, error) {
	response, err := c.httpClient.Get(fmt.Sprintf("%s/products/%s", c.url, sku))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("server returned %s", response.Status)
	}

	var product dto.Product

	if err := json.NewDecoder(response.Body).Decode(&product); err != nil {
		return false, err
	}

	fmt.Printf("Product: %+v\n", product)

	return true, nil
}
