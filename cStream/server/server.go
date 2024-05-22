package main

import (
	proto "cStream/protoc"
	"fmt"
	"io"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type srvr struct {
	proto.UnimplementedClientStreamServer
}

// hover over proto.ClientStream_ServerReplyServer **
func (s *srvr) ServerReply(strm proto.ClientStream_ServerReplyServer) error {
	total := 0 // keeps track of number of msgs from client

	for { // loop until EOF
		request, err := strm.Recv()
		if err == io.EOF {
			return strm.SendAndClose(&proto.SReply{ // sends response & closes call
				Resp: strconv.Itoa(total), // total is int >> int to string
			})
		}
		if err != nil {
			return err
		}

		total++
		fmt.Println(request)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterClientStreamServer(srv, &srvr{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
