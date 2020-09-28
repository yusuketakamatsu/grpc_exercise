package main

import (
	"context"
	"fmt"

	hello "src/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	client := hello.NewHelloClient(conn)
	message := &hello.HelloRequest{Name: "hello"}
	res, err := client.SayHello(context.Background(), message)
	fmt.Printf("%s\n", res)
}
