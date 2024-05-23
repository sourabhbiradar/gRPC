package main

import (
	proto "BiDir/protoc"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedBiDirServer
}

func (s *server) ServerReply(strm proto.BiDir_ServerReplyServer) error {
	for i := 0; i < 10; i++ { // sending
		err := strm.Send(&proto.SResp{Resp: "message" + strconv.Itoa(i) + "from server"})
		if err != nil {
			return errors.New("unable to send response")
		}
	}

	for { // receiving
		req, err := strm.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(req.Reqt)
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterBiDirServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
