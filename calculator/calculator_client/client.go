package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Client launched")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	fmt.Println("Starting to do a unary RPC")

	req := &calculatorpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Sum: %v", res.Result)
}
