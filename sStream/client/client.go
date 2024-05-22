package main

import (
	proto "SStream/protoc"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.SStreamClient

func main() {
	conn, err := grpc.NewClient(":2000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client = proto.NewSStreamClient(conn)

	r := gin.Default()
	r.GET("/sent", clientConnectionServer)
	r.Run(":8000")
}

func clientConnectionServer(c *gin.Context) {
	strm, err := client.ServerReply(context.TODO(), &proto.CReqt{Reqt: "requesting stream of replies"}) // sending context & request to server side

	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0 // count of replies

	for { // becoz we dont know how many replies are there
		message, err := strm.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("replies :", message)
		time.Sleep(1 * time.Second)
		count++
	}
	c.JSON(200, gin.H{
		"message-count": count,
	})
}
