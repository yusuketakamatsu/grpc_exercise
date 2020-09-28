package main

import (
	"context"
	"net"
	hello "src/pb"

	"google.golang.org/grpc"
)

type helloService struct{}

func (s *helloService) SayHello(ctx context.Context, message *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{
		Message: "hello",
	}, nil
}

func main() {
	port, err := net.Listen("tcp", ":3000")
	if err != nil {
		return
	}
	s := grpc.NewServer()

	hello.RegisterHelloServer(s, &helloService{})

	s.Serve(port)
}
