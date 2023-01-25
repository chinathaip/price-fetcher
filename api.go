package main

import (
	"context"
	"net/http"

	"github.com/chinathaip/price-fetcher/types"
	"github.com/labstack/echo/v4"
)

type JSONAPIServer struct {
	addr string
	svc  PriceFetcher
}

func NewJSONAPIServer(addr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{addr: addr, svc: svc}
}

func (server *JSONAPIServer) Run() {
	e := echo.New()
	e.GET("/", server.handleFetchPrice)
	e.Start(server.addr)
}

func (server *JSONAPIServer) handleFetchPrice(c echo.Context) error {
	coin := c.QueryParam("coin")
	if coin == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing coin param")
	}

	price, err := server.svc.FetchPrice(context.Background(), coin)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	resp := types.CoinPriceResponse{Name: coin, Price: price}

	return c.JSON(http.StatusOK, resp)
}
