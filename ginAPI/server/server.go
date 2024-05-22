package main

import (
	"context"
	"fmt"
	proto "ginAPI/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedWithAPIServer
}

func (s *server) ServerResp(c context.Context, req *proto.Creqt) (*proto.Sresp, error) {
	fmt.Println("received request from client :", req.Reqt)
	fmt.Println("Hello from server")

	return &proto.Sresp{}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterWithAPIServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
