package main

import (
	proto "SStream/protoc"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedSStreamServer
}

func (s *server) ServerReply(req *proto.CReqt, strm proto.SStream_ServerReplyServer) error {
	fmt.Println(req.Reqt) // printing client request
	time.Sleep(5 * time.Second)

	reply := []*proto.SResps{
		{Resps: "response 1"},
		{Resps: "response 1.5"},
		{Resps: "response 2"},
		{Resps: "response 3"},
	}

	for _, msg := range reply { // extracting one by one reply
		err := strm.Send(msg) // sending them to client as response
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":2000")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterSStreamServer(srv, &server{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
