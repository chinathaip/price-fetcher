package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcherService struct{}

func (s *priceFetcherService) FetchPrice(ctx context.Context, coin string) (float64, error) {
	return mockGetPrice(coin)
}

var mockPrice = map[string]float64{
	"BTC": 200_000.00,
	"ETH": 100_000.00,
}

func mockGetPrice(coin string) (float64, error) {
	price, ok := mockPrice[coin]
	if !ok {
		return price, fmt.Errorf("Not found")
	}
	return price, nil
}
