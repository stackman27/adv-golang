package main

import (
	"fmt"
	"log"
	"net"

	"github.com/go-microservice3/server"
	"github.com/go-microservice3/types"
	"google.golang.org/grpc"
)

 
func main() { 
	listen, err := net.Listen("tcp", ":9000") 
	if err != nil {
		log.Fatal("failed to listedn %v", err) 
		return
	}

	grpcServer := grpc.NewServer()
	types.RegisterMachineServer(grpcServer, &server.MachineServer{})

	err = grpcServer.Serve(listen) 	
	if err != nil {
		log.Fatal("failed to start server %v", err) 
		return
	}

	fmt.Printf("initializing gRPC server on port :9000")

}