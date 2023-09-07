package main

import (
	"context"
	"fmt"
)

type MetricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &MetricService{
		next: next,
	}
}


func (s *MetricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {	
	// your metric storage. Push to prometheus or something 

	fmt.Println("PUSHING METRIC SERVICES")
	return  s.next.FetchPrice(ctx, ticker)
}