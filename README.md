# go-grpc

This is a gRPC example project that demonstrates the implementation of user authentication and order management services using Go and gRPC.

## Project Structure

```
├── client
│   ├── order
│   │   └── main.go    # Order service client
│   └── user
│       └── main.go    # User service client
├── proto
│   ├── order
│   │   ├── order_grpc.pb.go
│   │   ├── order.pb.go
│   │   └── order.proto
│   └── user
│       ├── user_grpc.pb.go
│       ├── user.pb.go
│       └── user.proto
└── server
    ├── order
    │   └── order.go   # Order service implementation
    └── user
        └── user.go    # User service implementation
```

## Prerequisites

- Go 1.25 or later
- Protocol Buffers compiler (protoc)
- Go plugins for Protocol Buffers
  - `google.golang.org/protobuf/cmd/protoc-gen-go`
  - `google.golang.org/grpc/cmd/protoc-gen-go-grpc`

## Installation

1. Install the Protocol Buffers compiler:
   ```bash
   # For Ubuntu/Debian
   apt install -y protobuf-compiler
   
   # For macOS
   brew install protobuf
   ```

2. Install Go Protocol Buffers plugins:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. Clone the repository:
   ```bash
   git clone https://github.com/vsjadeja/go-grpc.git
   cd go-grpc
   ```

4. Generate gRPC code:
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/user/user.proto
   protoc --go_out=. --go-grpc_out=. proto/order/order.proto
   ```

## Running the Services

1. Start the user service:
   ```bash
   go run server/user/main.go
   ```

2. Start the order service:
   ```bash
   go run server/order/main.go
   ```

3. Run the user client:
   ```bash
   go run client/user/main.go
   ```

4. Run the order client:
   ```bash
   go run client/order/main.go
   ```

## Services

### User Service
The user service provides authentication and registration functionality:
- Login: Authenticates a user and returns a JWT token
- Register: Creates a new user account

### Order Service
The order service handles order management:
- Create Order: Creates a new order
- Get Order: Retrieves order details
- List Orders: Lists all orders for a user

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.