package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"grpctutorial/tutorial"
)

func main() {
	fmt.Println("gRPC client tutorial in Go")

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := tutorial.NewTutorialClient(conn)

	var yourName string
	fmt.Println("Please enter your name:")
	fmt.Scanf("%s", &yourName)
	request := &tutorial.HelloRequest{
		Name: yourName,
	}

	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println("Response:", response.Message)
}
