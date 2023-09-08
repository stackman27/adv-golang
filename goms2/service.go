package main

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcher is an interface that fetches the price.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)  
}

// pricefetcher implements the PriceFetcher implements
type pricefetcher struct {} 

// interface function
func (s *pricefetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string] float64 {
	"BTC": 20_000, 
	"ETH": 200,
	"GG": 100_000,
}	

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(100 * time.Millisecond)

	// http round trip to mimic this behavior 
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("Given ticker is not supported %s", ticker)
	}

	return price, nil
} 