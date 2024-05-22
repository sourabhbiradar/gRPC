package main

import (
	"context"
	proto "proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if any credentials pass in `NewCredentials()`

	if err != nil {
		panic(err)
	}

	client := proto.NewExampleClient(conn) // registering client

	// preparing grpc request
	req := &proto.HelloReqt{Hreqt: "hello from client"}

	// sending request
	client.ServerReply(context.TODO(), req)

}
