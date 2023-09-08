package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/go-microservice/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc PriceFetcher
}

func NewJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc: svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(fn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestId", rand.Intn(1000000))

	// this is where all the errros are getting handled
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker) 
	if err != nil {
		return err 
	}

	priceResp := &types.PriceResponse{
		Ticker: ticker,
		Price: price,
	}

	return writeJSON(w, http.StatusOK, priceResp)
}          


func writeJSON(w http.ResponseWriter, s int, v any) error  {
	w.WriteHeader(s) 
	return json.NewEncoder(w).Encode(v) 
}