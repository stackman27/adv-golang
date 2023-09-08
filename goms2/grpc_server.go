package main

import (
	"context"
	"math/rand"
	"net"

	"github.com/go-microservice/types"
	"google.golang.org/grpc"
)

func makeGRPCServerAndRun(listenAddr string, svc PriceFetcher) error {
	grpcPriceFetcher := NewGRPCPriceFetcher(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	types.RegisterPriceFetcherServer(server, grpcPriceFetcher)

	return server.Serve(ln)
}

type GRPCPriceFetcherServer struct {
	svc PriceFetcher
	types.UnimplementedPriceFetcherServer // combining interface from types.pb.go file
} 

func NewGRPCPriceFetcher(svc PriceFetcher) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *types.PriceRequest) (*types.PriceResponse, error) {
	reqId := rand.Intn(2000)
	ctx = context.WithValue(ctx, "requestId", reqId)

	resp, err := s.svc.FetchPrice(ctx, req.GetToken())
	if err != nil {
		return nil, err
	}

	return &types.PriceResponse{
		Ticker: req.GetToken(),
		Price:  float32(resp),
	}, nil

}
