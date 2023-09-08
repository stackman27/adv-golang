package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-microservice/client"
	"github.com/go-microservice/types"
)

func main() {
	var (
		jsonPortAddr = flag.String("jsonPortAddr", ":3000", "listen address of the json transport")
		grpcPortAddr = flag.String("grpc", ":4000", "listen address of the grpc transport") 
	 
		ctx = context.Background()
	)
	
	flag.Parse() 
 	svc := NewLoginService(NewMetricService(&pricefetcher{}))

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err) 
	}

	go func() {
		time.Sleep(3 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &types.PriceRequest{
			Token: "BTC",
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", resp)
	}()

	go makeGRPCServerAndRun(*grpcPortAddr, svc)  

	jsonServer := NewJSONAPIServer(*jsonPortAddr, svc)	
	jsonServer.Run()
}