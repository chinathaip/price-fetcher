package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chinathaip/price-fetcher/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (c *Client) FetchPrice(ctx context.Context, coin string) (*types.CoinPriceResponse, error) {
	url := fmt.Sprintf("%s?coin=%s", c.endpoint, coin)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var priceResponse types.CoinPriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		return nil, err
	}

	return &priceResponse, nil
}
