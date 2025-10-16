# Stage 1: Build
FROM golang:1.25-alpine AS builder

# Install protoc-gen-go if needed
RUN apk add --no-cache git bash protoc build-base

WORKDIR /app

# Install protoc-gen-go if needed
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Cache go modules
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Generate Go code from .proto files
RUN protoc --go_out=. --go-grpc_out=. proto/order/order.proto

# Get build metadata
ARG VERSION
ARG COMMIT
ARG BUILD_DATE

# Build the gRPC server
RUN go build  -ldflags="-s -w \
    -X 'main.Version=${VERSION}' \
    -X 'main.Commit=${COMMIT}' \
    -X 'main.BuildDate=${BUILD_DATE}'" \
    -o order-server server/order/order.go

# Stage 2: Run
FROM alpine:3.18

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/order-server .

# Expose gRPC port
EXPOSE 8080

# Run server
CMD ["./order-server"]