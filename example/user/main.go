package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-grpc/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const userServicePort string = "8080"

func main() {
	conn, err := grpc.NewClient("user-server:"+userServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Login(ctx, &pb.LoginRequest{UserName: "Virendra", Password: "password"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}

	fmt.Println("Token:", res.Token)
}
