package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-microservice/types"
	"google.golang.org/grpc"
)
 

func NewGRPCClient(remoteAddr string) (types.PriceFetcherClient, error){
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := types.NewPriceFetcherClient(conn)

	return c, nil
}

type client struct {
	endpoint string
}

func New(endpoint string) *client {
	return &client{
		endpoint: endpoint,
	}
}

func (c *client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker) 

	req, err := http.NewRequest("get", endpoint, nil) 
	if err != nil {
		return nil, err 
	}

	resp, err := http.DefaultClient.Do(req) 
	if err != nil {
		return nil, err 
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}

		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err 
		}

		return nil, fmt.Errorf("service responded without 200 OK response %s", httpErr["error"])
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err 
	}

	return priceResp, nil
}