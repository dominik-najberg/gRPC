package main

import (
	"context"
	"fmt"
	"gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("This is the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // we have no certificate
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close() // this will be called at the end of the program

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a unary RPC")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dominik",
			LastName:  "Najberg",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Greeting: %v", res.Result)
}
