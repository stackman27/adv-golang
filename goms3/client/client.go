package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/go-microservice3/types"
	"google.golang.org/grpc"
)


func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Issue establishing connection with the server")
		return 
	}

	defer conn.Close()
	client := types.NewMachineClient(conn)
	
	// Execute 
	instructions := []*types.Instruction{
		{Operator: "PUSH", Operand: 100},
		{Operator: "PUSH", Operand: 2},
		{Operator: "ADD"},
	}

	RunExecute(client, instructions)
}

func RunExecute(client types.MachineClient, instructions []*types.Instruction) {
	log.Printf("Streaming %v", instructions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 

	stream, err := client.Execute(ctx) 
	if err != nil {
		log.Fatalf("%v.Execute(ctx) = %v, %v: ", client, stream, err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF")
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Err: %v", err)
			}
			log.Printf("output: %v", result.GetOutput())
		}
	}()


	for _, instruction := range instructions {
		err := stream.Send(instruction)
		if err != nil {
			log.Fatalf("%v.Send(%v) = %v: ", stream, instruction, err)
		}

		time.Sleep(500 * time.Millisecond)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("%v.CloseSend() got error %v, want %v", stream, err, nil)
	}

	<-waitc
}