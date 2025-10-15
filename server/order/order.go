package main

import (
	"context"
	"fmt"
	pb "go-grpc/proto/order"
	"log"
	"net"

	"google.golang.org/grpc"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
}

func (s *OrderServer) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	// Implement order listing logic here
	return &pb.ListResponse{Orders: []*pb.Order{
		{Id: "order-1", ItemName: "user-1", Amount: 2},
		{Id: "order-2", ItemName: "user-2", Amount: 1},
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &OrderServer{})

	fmt.Println("gRPC OrderServer running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
