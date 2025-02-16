package main

import (
	"context"
	"log"
	"net"
	"sync"

	// Import the generated protobuf code
	pb "grpc-counter/proto/gen/go"

	"google.golang.org/grpc"
)

// Define the server struct
type server struct {
	pb.UnimplementedCounterServiceServer
	mu      sync.Mutex
	counter int64
}

// Increment RPC implementation
func (s *server) Increment(ctx context.Context, req *pb.IncrementRequest) (*pb.IncrementResponse, error) {
	s.mu.Lock()
	s.counter++
	value := s.counter
	s.mu.Unlock()
	return &pb.IncrementResponse{Value: value}, nil
}

// Main function to start the gRPC server
func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	pb.RegisterCounterServiceServer(s, &server{})

	// Log the server running status
	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
