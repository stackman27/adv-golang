package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// pricefetcher implements the PriceFetcher implements
type loginService struct {
	next PriceFetcher
}

func NewLoginService(next PriceFetcher) PriceFetcher {
	return &loginService{
		next: next,
	}
}

func (s *loginService) FetchPrice(ctx context.Context, ticker string) (price float64, err error){
	defer func(begin time.Time) { 
		logrus.WithFields(logrus.Fields{
			"requestId": ctx.Value("requestId"),
			"took": time.Since(begin),
			"err": err, 
			"price": price,
			"ticker": ticker, 
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}