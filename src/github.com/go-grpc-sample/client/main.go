package main

import (
	"context"
	"fmt"
	greeter "github.com/go-grpc-sample/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("127.0.0.1:55555", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect: %s", err)
	}
	defer connection.Close()

	client := greeter.NewGreeterClient(connection)

	context, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	req := &greeter.HelloRequest{
		Name: "hghyk023",
	}

	rep, err := client.SayHello(context, req)
	fmt.Println(rep.Message)

}
