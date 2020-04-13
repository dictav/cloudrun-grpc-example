package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/dictav/cloudrun-grpc-example/proto"

	"google.golang.org/grpc"
)

// Prove that server implements proto.GreeterServer
var _ proto.GreeterServer = (*server)(nil)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ep := ":"+port
	s := grpc.NewServer()

	proto.RegisterGreeterServer(s, newServer())

	listen, err := net.Listen("tcp", ep)
	if err != nil {
		os.Exit(2)
	}

	println("Starting: gRPC Listener", ep)
	if err := s.Serve(listen); err != nil {
		os.Exit(2)
	}
}


// server is a struct implements the proto.GreeterServer
type server struct {
}

// New returns a new Server
func newServer() *server {
	return &server{}
}

// Greet sys hello
func (s *server) Greet(ctx context.Context, r *proto.GreetRequest) (*proto.GreetResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, fmt.Errorf("client cancelled: abandoning")
	}

	greeting := r.GetType().String()

	return &proto.GreetResponse{
		Greeting: greeting + ", " + r.GetName(),
	}, nil
}
