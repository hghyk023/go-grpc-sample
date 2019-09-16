package main

import (
	"context"
	"github.com/go-grpc-sample/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GreeterService struct {
}

func (s *GreeterService) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{
		Message: "Hello " + in.Name,
	}, nil

}

func (s *GreeterService) SayHelloAgain(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:55555")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := &GreeterService{}

	greeter.RegisterGreeterServer(server, service)

	server.Serve(listen)

}
