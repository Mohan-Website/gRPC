package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
	"google.golang.org/grpc"
)

// Server implements the chat.ChatServiceServer interface
type Server struct {
	chat.UnimplementedChatServiceServer
}

// SayHello implements the SayHello RPC method from the ChatServiceServer interface
func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	return &chat.Message{Body: "Hello, " + in.Body}, nil
}

// BroadcastMessage implements the BroadcastMessage RPC method from the ChatServiceServer interface
func (s *Server) BroadcastMessage(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	// Implement the logic for broadcasting messages here
	return &chat.Message{Body: "Broadcast: " + in.Body}, nil
}

func main() {
	fmt.Println("Server Started")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	chat.RegisterChatServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
