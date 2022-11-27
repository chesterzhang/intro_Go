package __

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *ChatMessage ) (*ChatMessage, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &ChatMessage{Body: "Hello From the Server!"}, nil
}

