package main

import (
	proto "BiDir/protoc"
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.BiDirClient

func main() {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client = proto.NewBiDirClient(conn)

	r := gin.Default()
	r.GET("/sent", clientConnServer)
	r.Run(":6000")
}

func clientConnServer(c *gin.Context) {
	strm, err := client.ServerReply(context.TODO())

	if err != nil {
		fmt.Println("somethings wrong")
		return
	}

	send, receive := 0, 0

	for i := 0; i < 10; i++ {
		err := strm.Send(&proto.CReq{Reqt: "msg" + strconv.Itoa(i) + "from client"})
		if err != nil {
			fmt.Println("unable to send data")
			return
		}
		send++
	}
	if err := strm.CloseSend(); err != nil {
		log.Println(err)
	}

	for {
		msg, err := strm.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("server message:", msg)
		receive++
	}
	c.JSON(200, gin.H{
		"messages sent :":     send,
		"messages received :": receive,
	})
}
