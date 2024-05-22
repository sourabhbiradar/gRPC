package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	proto "proto" // alias

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func (s *server) ServerReply(c context.Context, req *proto.HelloReqt) (*proto.HelloResp, error) {
	fmt.Println("receive request", req.Hreqt)
	fmt.Println("server started ...")

	return &proto.HelloResp{}, errors.New("")
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")

	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer() // engine

	proto.RegisterExampleServer(srv, &server{})

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}

}
