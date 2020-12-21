package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GeorgHs/GoServer/ServerClient/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Println("Created client: %f",c)

	doUnary(c)
	//c.Greet(context.Background(), in*greetpb.GreetRequest)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting to do a Unary RPC")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.SumResult)
}
