package httpClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/internal/service/dto"
)

type InventoryClient interface {
	CheckAvailability(sku string, quantity int) (bool, error)
	ReserveProduct(sku string, quantity int) error
	CancelOrder(sku string, quantity int) error
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

func (client *InventoryClientImpl) CheckAvailability(sku string, quantity int) (bool, error) {
	response, err := client.httpClient.Get(fmt.Sprintf("%s/products/%s", client.url, sku))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("inventory-service returned http code: %s", response.Status)
	}

	var product dto.Product

	if err := json.NewDecoder(response.Body).Decode(&product); err != nil {
		return false, fmt.Errorf("error decoding response from inventory-service: %s", err)
	}

	return product.Quantity >= quantity, nil
}

func (client *InventoryClientImpl) ReserveProduct(sku string, quantity int) error {
	return nil
}

func (client *InventoryClientImpl) CancelOrder(sku string, quantity int) error {
	return nil
}
