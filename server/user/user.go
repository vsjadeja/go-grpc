package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "go-grpc/proto/user"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
)

const JWT_SECRET = "my-jwt-secret"

var jwtSecret = []byte(JWT_SECRET)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received login request for user: %s", req.GetUserName())
	token, err := GenerateJWT(req.GetUserName())
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

func (s *UserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received register request for user: %s", req.GetName())
	return &pb.RegisterResponse{UserId: "some-user-id"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserServer{})

	fmt.Println("gRPC UserServer running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GenerateJWT(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // 1 hour expiry
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
