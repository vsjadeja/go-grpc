package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-grpc/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.List(ctx, &pb.ListRequest{UserId: "user-1"})
	if err != nil {
		log.Fatalf("could not list orders: %v", err)
	}

	fmt.Println("Orders:", res.Orders)
}
