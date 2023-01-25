package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggerService struct {
	next PriceFetcher
}

func NewLoggerService(next PriceFetcher) PriceFetcher {
	return &loggerService{
		next: next,
	}
}

func (s *loggerService) FetchPrice(ctx context.Context, coin string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(begin),
			"price": price,
			"err":   err,
		}).Info("fetchPrice")
	}(time.Now())
	return s.next.FetchPrice(ctx, coin)
}
