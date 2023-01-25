package main

import (
	"context"
	"log"
)

func main() {
	svc := NewLoggerService(&priceFetcherService{})

	price, err := svc.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatalf("error : %v", err.Error())
	}
	log.Println(price)
}
