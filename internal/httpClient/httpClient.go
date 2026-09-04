package httpClient

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	response, err := c.httpClient.Get(fmt.Sprintf(""))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("server returned %s", response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(&sku); err != nil {
		return false, err
	}

	return true, nil
}
